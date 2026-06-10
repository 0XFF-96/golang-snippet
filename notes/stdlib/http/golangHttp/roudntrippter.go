package golangHttp

type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
}

// RoundTrip implements the RoundTripper interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see Get, Post, and the Client type.
//
// Like the RoundTripper interface, the error types returned
// by RoundTrip are unspecified.
func (t *Transport) RoundTrip(req *Request) (*Response, error) {
	return t.roundTrip(req)
}

func (t *Transport) roundTrip(req *Request) (*Response, error) {
	// t.nextProtoOnce.Do(t.onceSetNextProtoDefaults)
	ctx := req.Context()

	// 这里面是 validate req 的地方
	//1. req.URL
	//2. if req.Header
	//3. scheme := req.URL.Scheme
	//4. isHTTP := scheme == "http" || scheme == "https"
	//5. if t.useRegisteredProtocol(req) {
	//6. req.Method != "" && !validMethod(req.Method) {
	//7. req.URL.Host == "" {

	for {
		select {
		case <-ctx.Done():
			req.closeBody()
			return nil, ctx.Err()
		default:
		}

		// treq gets modified by roundTrip, so we need to recreate for each retry.
		treq := &transportRequest{Request: req, trace: trace}
		cm, err := t.connectMethodForRequest(treq)
		if err != nil {
			req.closeBody()
			return nil, err
		}

		// Get the cached or newly-created connection to either the
		// host (for http or https), the http proxy, or the http proxy
		// pre-CONNECTed to https server. In any case, we'll be ready
		// to send it requests.
		pconn, err := t.getConn(treq, cm)
		if err != nil {
			t.setReqCanceler(req, nil)
			req.closeBody()
			return nil, err
		}

		var resp *Response
		if pconn.alt != nil {
			// HTTP/2 path.
			t.setReqCanceler(req, nil) // not cancelable with CancelRequest
			resp, err = pconn.alt.RoundTrip(req)
		} else {
			resp, err = pconn.roundTrip(treq)
		}
		if err == nil {
			return resp, nil
		}

		// Failed. Clean up and determine whether to retry.
	}
}

// 底层是由 presistConn 的 roundTrip 真正实现的


// queueForDial queues w to wait for permission to begin dialing.
// Once w receives permission to dial, it will do so in a separate goroutine.
func (t *Transport) queueForDial(w *wantConn) {
	w.beforeDial()

	// 没有连接的情况下直接 dial
	/*if t.MaxConnsPerHost <= 0 {
		go t.dialConnFor(w)
		return
	}*/

	t.connsPerHostMu.Lock()
	defer t.connsPerHostMu.Unlock()

	if n := t.connsPerHost[w.key]; n < t.MaxConnsPerHost {
		if t.connsPerHost == nil {
			t.connsPerHost = make(map[connectMethodKey]int)
		}
		t.connsPerHost[w.key] = n + 1
		go t.dialConnFor(w)
		return
	}

	// 这里直接返回了
	// 谁异步 dial conn 给 wantConnQueue ?
	if t.connsPerHostWait == nil {
		t.connsPerHostWait = make(map[connectMethodKey]wantConnQueue)
	}
	q := t.connsPerHostWait[w.key]
	q.cleanFront()
	q.pushBack(w)
	t.connsPerHostWait[w.key] = q
}

// queueForIdleConn queues w to receive the next idle connection for w.cm.
// As an optimization hint to the caller, queueForIdleConn reports whether
// it successfully delivered an already-idle connection.
func (t *Transport) queueForIdleConn(w *wantConn) (delivered bool) {
	//if t.DisableKeepAlives {
	//	return false
	//}

	t.idleMu.Lock()
	defer t.idleMu.Unlock()

	// Stop closing connections that become idle - we might want one.
	// (That is, undo the effect of t.CloseIdleConnections.)
	t.closeIdle = false

	if w == nil {
		// Happens in test hook.
		return false
	}

	// Look for most recently-used idle connection.
	if list, ok := t.idleConn[w.key]; ok {
		stop := false
		delivered := false
		for len(list) > 0 && !stop {
			pconn := list[len(list)-1]
			if pconn.isBroken() {
				// persistConn.readLoop has marked the connection broken,
				// but Transport.removeIdleConn has not yet removed it from the idle list.
				// Drop on floor on behalf of Transport.removeIdleConn.
				list = list[:len(list)-1]
				continue
			}
			delivered = w.tryDeliver(pconn, nil)
			if delivered {
				if pconn.alt != nil {
					// HTTP/2: multiple clients can share pconn.
					// Leave it in the list.
				} else {
					// HTTP/1: only one client can use pconn.
					// Remove it from the list.
					// t.idleLRU.remove(pconn)
					list = list[:len(list)-1]
				}
			}
			stop = true
		}
		if len(list) > 0 {
			t.idleConn[w.key] = list
		} else {
			delete(t.idleConn, w.key)
		}
		if stop {
			return delivered
		}
	}

	// Register to receive next connection that becomes idle.
	if t.idleConnWait == nil {
		t.idleConnWait = make(map[connectMethodKey]wantConnQueue)
	}
	q := t.idleConnWait[w.key]

	// q.cleanFront()
	// q.pushBack(w)
	t.idleConnWait[w.key] = q
	return false
}