package golangHttp

import (
	"sync"
)

type requestAndChan struct {
	req *Request
	callerGone <-chan struct{} // closed when roundTrip caller has returned
}

type writeRequest struct {
	// req *transportRequest
	ch  chan<- error

	// Optional blocking chan for Expect: 100-continue (for receive).
	// If not nil, writeLoop blocks sending request body until
	// it receives from this chan.
	continueCh <-chan struct{}
}


type persistConn struct {
	// alt optionally specifies the TLS NextProto RoundTripper.
	// This is used for HTTP/2 today and future protocols later.
	// If it's non-nil, the rest of the fields are unused.
	alt RoundTripper

	reqch     chan requestAndChan // written by roundTrip; read by readLoop
	writech   chan writeRequest   // written by roundTrip; read by writeLoop

	t         *Transport
	closech   chan struct{}       // closed when conn closed
	writeLoopDone chan struct{} // closed when write loop ends

	mu                   sync.Mutex // guards following fields
}

func (pc *persistConn) readLoop() {
}

func (pc *persistConn) writeLoop() {
	defer close(pc.writeLoopDone)
	for {
		select {
		case <-pc.closech:
			return
		}
	}
}

// isBroken reports whether this connection is in a known broken state.
func (pc *persistConn) isBroken() bool {
	pc.mu.Lock()
	b := pc.closed != nil
	pc.mu.Unlock()
	return b
}

// A wantConnQueue is a queue of wantConns.
type wantConnQueue struct {
	// This is a queue, not a deque.
	// It is split into two stages - head[headPos:] and tail.
	// popFront is trivial (headPos++) on the first stage, and
	// pushBack is trivial (append) on the second stage.
	// If the first stage is empty, popFront can swap the
	// first and second stages to remedy the situation.
	//
	// This two-stage split is analogous to the use of two lists
	// in Okasaki's purely functional queue but without the
	// overhead of reversing the list when swapping stages.
	head    []*wantConn
	headPos int
	tail    []*wantConn
}