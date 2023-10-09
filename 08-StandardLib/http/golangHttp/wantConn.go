package golangHttp

import (
	"context"
	"sync"
)

// A wantConn records state about a wanted connection
// (that is, an active call to getConn).
// The conn may be gotten by dialing or by finding an idle connection,
// or a cancellation may make the conn no longer wanted.
// These three options are racing against each other and use
// wantConn to coordinate and agree about the winning outcome.
type wantConn struct {
	cm    connectMethod
	key   connectMethodKey // cm.key()
	ctx   context.Context  // context for dial
	ready chan struct{}    // closed when pc, err pair is delivered

	// hooks for testing to know when dials are done
	// beforeDial is called in the getConn goroutine when the dial is queued.
	// afterDial is called when the dial is completed or cancelled.
	beforeDial func()
	afterDial  func()

	mu  sync.Mutex // protects pc, err, close(ready)
	pc  *persistConn
	err error
}


// wantConn 是一个桥梁， 连接 Transport 和  persistConn
// tryDeliver attempts to deliver pc, err to
// w and reports whether it succeeded.
func (w *wantConn) tryDeliver(pc *persistConn, err error) bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.pc != nil || w.err != nil {
		return false
	}
	w.pc = pc
	w.err = err
	if w.pc == nil && w.err == nil {
		panic("net/http: internal error: misuse of tryDeliver")
	}
	close(w.ready)
	return true
}



// 核心方法2
// func (t *Transport) dialConn(ctx context.Context, cm connectMethod) (pconn *persistConn, err error) {
//	pconn = &persistConn{

// 核心方法3
// Get the cached or newly-created connection to either the
// host (for http or https), the http proxy, or the http proxy
// pre-CONNECTed to https server. In any case, we'll be ready
// to send it requests.
// pconn, err := t.getConn(treq, cm)