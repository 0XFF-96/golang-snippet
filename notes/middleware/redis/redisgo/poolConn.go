package redisgo

type poolConn struct {
	c          Conn
	t          time.Time
	created    time.Time
	next, prev *poolConn
}

type Pool struct {
	// Dial is an application supplied function for creating and configuring a
	// connection.
	//
	// The connection returned from Dial must not be in a special state
	// (subscribed to pubsub channel-antomay, transaction started, ...).
	Dial func() (Conn, error)

	// DialContext is an application supplied function for creating and configuring a
	// connection with the given context.
	//
	// The connection returned from Dial must not be in a special state
	// (subscribed to pubsub channel-antomay, transaction started, ...).
	DialContext func(ctx context.Context) (Conn, error)

	// TestOnBorrow is an optional application supplied function for checking
	// the health of an idle connection before the connection is used again by
	// the application. Argument t is the time that the connection was returned
	// to the pool. If the function returns an error, then the connection is
	// closed.
	TestOnBorrow func(c Conn, t time.Time) error

	// Maximum number of idle connections in the pool.
	MaxIdle int

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration

	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool

	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime time.Duration

	chInitialized uint32 // set to 1 when field ch is initialized

	mu           sync.Mutex    // mu protects the following fields
	closed       bool          // set to true when the pool is closed.
	active       int           // the number of open connections in the pool
	ch           chan struct{} // limits open connections when p.Wait is true
	idle         idleList      // idle connections
	waitCount    int64         // total number of connections waited for.
	waitDuration time.Duration // total time waited for new connections.
}

// 为什么需要 lazyInit ?
func (p *Pool) lazyInit() {
	// Fast path.
	if atomic.LoadUint32(&p.chInitialized) == 1 {
		return
	}
	// Slow path.
	p.mu.Lock()
	if p.chInitialized == 0 {
		p.ch = make(chan struct{}, p.MaxActive)
		if p.closed {
			close(p.ch)
		} else {
			for i := 0; i < p.MaxActive; i++ {
				p.ch <- struct{}{}
			}
		}
		atomic.StoreUint32(&p.chInitialized, 1)
	}
	p.mu.Unlock()
}


// get prunes stale connections and returns a connection from the idle list or
// creates a new connection.
func (p *Pool) get(ctx context.Context) (*poolConn, error) {

	// Handle limit for p.Wait == true.
	var waited time.Duration
	if p.Wait && p.MaxActive > 0 {
		p.lazyInit()

		// wait indicates if we believe it will block so its not 100% accurate
		// however for stats it should be good enough.
		wait := len(p.ch) == 0
		var start time.Time
		if wait {
			start = time.Now()
		}
		if ctx == nil {
			<-p.ch
		} else {
			select {
			case <-p.ch:
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		if wait {
			waited = time.Since(start)
		}
	}

	p.mu.Lock()

	if waited > 0 {
		p.waitCount++
		p.waitDuration += waited
	}

	// Prune stale connections at the back of the idle list.
	if p.IdleTimeout > 0 {
		n := p.idle.count
		for i := 0; i < n && p.idle.back != nil && p.idle.back.t.Add(p.IdleTimeout).Before(nowFunc()); i++ {
			pc := p.idle.back
			p.idle.popBack()
			p.mu.Unlock()
			pc.c.Close()
			p.mu.Lock()
			p.active--
		}
	}

	// Get idle connection from the front of idle list.
	for p.idle.front != nil {
		pc := p.idle.front
		p.idle.popFront()
		p.mu.Unlock()
		if (p.TestOnBorrow == nil || p.TestOnBorrow(pc.c, pc.t) == nil) &&
			(p.MaxConnLifetime == 0 || nowFunc().Sub(pc.created) < p.MaxConnLifetime) {
			return pc, nil
		}
		pc.c.Close()
		p.mu.Lock()
		p.active--
	}

	// Check for pool closed before dialing a new connection.
	if p.closed {
		p.mu.Unlock()
		return nil, errors.New("redigo: get on closed pool")
	}

	// Handle limit for p.Wait == false.
	if !p.Wait && p.MaxActive > 0 && p.active >= p.MaxActive {
		p.mu.Unlock()
		return nil, ErrPoolExhausted
	}

	p.active++
	p.mu.Unlock()
	c, err := p.dial(ctx)
	if err != nil {
		c = nil
		p.mu.Lock()
		p.active--
		if p.ch != nil && !p.closed {
			p.ch <- struct{}{}
		}
		p.mu.Unlock()
	}
	return &poolConn{c: c, created: nowFunc()}, err
}

