package msyql

import (
	"sync"
	"time"

	"source-code-reading/go/src/database/sql/driver"
)

type Conn struct {
	db *DB

	// closemu prevents the connection from closing while there
	// is an active query. It is held for read during queries
	// and exclusively during close.
	closemu sync.RWMutex

	// dc is owned until close, at which point
	// it's returned to the connection pool.
	dc *driverConn

	// done transitions from 0 to 1 exactly once, on close.
	// Once done, all operations fail with ErrConnDone.
	// Use atomic operations on value when checking value.
	done int32
}



// DB is a database handle representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple
// goroutines.
//
// The sql package creates and frees connections automatically; it
// also maintains a free pool of idle connections.
// If the database has
// a concept of per-connection state, such state can be reliably observed
// within a transaction (Tx) or connection (Conn). Once DB.Begin is called, the
// returned Tx is bound to a single connection. Once Commit or
// Rollback is called on the transaction, that transaction's
// connection is returned to DB's idle connection pool. The pool size
// can be controlled with SetMaxIdleConns.
type DB struct {
	// Atomic access only. At top of struct to prevent mis-alignment
	// on 32-bit platforms. Of type time.Duration.
	// waitDuration int64 // Total time waited for new connections.

	connector driver.Connector
	// numClosed is an atomic counter which represents a total number of
	// closed connections. Stmt.openStmt checks it before cleaning closed
	// connections in Stmt.css.
	// numClosed uint64

	// mu           sync.Mutex // protects following fields
	// freeConn     []*driverConn


	// 这个 map 用来干嘛的？什么时候用的？
	connRequests map[uint64]chan connRequest
	resetterCh        chan *driverConn
	// nextRequest  uint64 // Next key to use in connRequests.
	// numOpen      int    // number of opened and pending open connections
	// Used to signal the need for new connections
	// a goroutine running connectionOpener() reads on this chan and
	// maybeOpenNewConnections sends on the chan (one send per needed connection)
	// It is closed during db.Close(). The close tells the connectionOpener
	// goroutine to exit.
	openerCh          chan struct{}
	closed            bool

	// 不知道这个是干嘛用的
	// dep               map[finalCloser]depSet
	lastPut           map[*driverConn]string // stacktrace of last conn's put; debug only
	//maxIdleCount      int                    // zero means defaultMaxIdleConns; negative means 0
	//maxOpen           int                    // <= 0 means unlimited
	//maxLifetime       time.Duration          // maximum amount of time a connection may be reused
	//maxIdleTime       time.Duration          // maximum amount of time a connection may be idle before being closed
	//cleanerCh         chan struct{}
	//waitCount         int64 // Total number of connections waited for.
	//maxIdleClosed     int64 // Total number of connections closed due to idle count.
	//maxIdleTimeClosed int64 // Total number of connections closed due to idle time.
	//maxLifetimeClosed int64 // Total number of connections closed due to max connection lifetime limit.

	stop func() // stop cancels the connection opener.
}

// driverConn wraps a driver.Conn with a mutex, to
// be held during all calls into the Conn. (including any calls onto
// interfaces returned via that Conn, such as calls on Tx, Stmt,
// Result, Rows)
type driverConn struct {
	db        *DB
	createdAt time.Time

	sync.Mutex  // guards following
	// ci          driver.Conn
	needReset   bool // The connection session should be reset before use if true.
	closed      bool
	finalClosed bool // ci.Close has been called
	// openStmt    map[*driverStmt]bool

	// guarded by db.mu
	inUse      bool
	returnedAt time.Time // Time the connection was created or returned.
	onPut      []func()  // code (with db.mu held) run when conn is next returned
	dbmuClosed bool      // same as closed, but guarded by db.mu, for removeClosedStmtLocked
}

// connRequest represents one request for a new connection
// When there are no idle connections available, DB.conn will create
// a new connRequest and put it on the db.connRequests list.
type connRequest struct {
	conn *driverConn
	err  error
}