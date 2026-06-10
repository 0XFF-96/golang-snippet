package golangHttp

import (
	"container/list"
	"context"
	"sync"
)

type Transport struct {
	idleMu       sync.Mutex
	closeIdle    bool                                // user has requested to close all idle conns

	// Slide down, keeping most recently-used
	// conns at the end.
	idleConn     map[connectMethodKey][]*persistConn // most recently used at end
	idleConnWait map[connectMethodKey]wantConnQueue  // waiting getConns

	// lru 是为了 idle 达到 max 时，找出最近未使用的一个连接？为何不找 idleConn 的开头一个？
	idleLRU      connLRU
}

type connLRU struct {
	ll *list.List // list.Element.Value type is of *persistConn
	m  map[*persistConn]*list.Element
}


// connectMethodKey is the map key version of connectMethod, with a
// stringified proxy URL (or the empty string) instead of a pointer to
// a URL.
type connectMethodKey struct {
	proxy, scheme, addr string
	onlyH1              bool
}


func (t *Transport) dialConn(ctx context.Context, cm connectMethod) (pconn *persistConn, err error) {
	// 1. 初始化 persistConn 的地方
	// 2. 初始化读 buffer 和 写 buffer 的地方
	pconn = &persistConn{
	}

	// 两个核心方法
	go pconn.readLoop()
	go pconn.writeLoop()
	return pconn, nil
}

func (t *Transport) dialConnFor(w *wantConn) {
	defer w.afterDial()

	pc, err := t.dialConn(w.ctx, w.cm)

	delivered := w.tryDeliver(pc, err)
	if err == nil && (!delivered || pc.alt != nil) {
		// pconn was not passed to w,
		// or it is HTTP/2 and can be shared.
		// Add to the idle connection pool.
		t.putOrCloseIdleConn(pc)
	}
}


func (t *Transport) putOrCloseIdleConn(pconn *persistConn) {
	//if err := t.tryPutIdleConn(pconn); err != nil {
	//	pconn.close(err)
	//}
}


func (t *Transport) getConn(treq *transportRequest, cm connectMethod) (pc *persistConn, err error) {
	req := treq.Request
	ctx := req.Context()


	// 1. 初始化 wantConn 的地方
	w := &wantConn{
		cm:         cm,
		key:        cm.key(),
		ctx:        ctx,
		ready:      make(chan struct{}, 1),

		// 生命周期钩子函数
		// beforeDial: testHookPrePendingDial,
		// afterDial:  testHookPostPendingDial,
	}

	// Queue for idle connection.
	// 看是否拿到 idle conn
	if delivered := t.queueForIdleConn(w); delivered {
		pc := w.pc
		// Trace only for HTTP/1.
		// HTTP/2 calls trace.GotConn itself.
		if pc.alt == nil && trace != nil && trace.GotConn != nil {
			trace.GotConn(pc.gotIdleConnTrace(pc.idleAt))
		}
		// set request canceler to some non-nil function so we
		// can detect whether it was cleared between now and when
		// we enter roundTrip
		t.setReqCanceler(req, func(error) {})
		return pc, nil
	}

	// 不太懂需要干嘛
	// cancelc := make(chan error, 1)
	// t.setReqCanceler(req, func(err error) { cancelc <- err })

	// 没有拿到 idle conn, 放在需要获取 idle conn 的池子
	// Queue for permission to dial.
	t.queueForDial(w)

	// Wait for completion or cancellation.
	// 如果缓冲池里面没有 conn 的时候
	// 需要等待
	return nil, nil
}


