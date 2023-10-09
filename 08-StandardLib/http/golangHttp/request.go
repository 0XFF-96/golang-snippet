package golangHttp

import (
	"context"
	"net/url"
)

type Request struct {
	// 1.
	URL *url.URL

	// 2.
	Method string

	// 3. 为什么 Request 需要内嵌一个 Context ?
	// ctx is either the client or server context. It should only
	// be modified via copying the whole Request using WithContext.
	// It is unexported to prevent people from using Context wrong
	// and mutating the contexts held by callers of the same request. ?
	ctx context.Context
}

func (r *Request) closeBody() {
}


// 这是为什么需要 Context 的原因
// 1. for outgoing request
// 2. for incoming request
// Context returns the request's context. To change the context, use
// WithContext.
//
// The returned context is always non-nil; it defaults to the
// background context.
//
// For outgoing client requests, the context controls cancellation.
//
// For incoming server requests, the context is canceled when the
// client's connection closes, the request is canceled (with HTTP/2),
// or when the ServeHTTP method returns.
func (r *Request) Context() context.Context {
	if r.ctx != nil {
		return r.ctx
	}
	return context.Background()
}