package golangHttp

import (
	"errors"
	"net/url"
	"time"
)

type Client struct {
	Transport RoundTripper
	CheckRedirect func(req *Request, via []*Request) error


	// 1. 什么情况 cooieJar 会等于 nil ?
	// 2. 在 send 方法的时候，为什么需要对偶设置 cookie jar
	// 3. 在请求时， req.AddCookie(cookie)
	// 4. 在有 response 时， c.Jar.SetCookies(req.URL, rc)
	Jar CookieJar
	Timeout time.Duration
}


func (c *Client) Do(req *Request) (*Response, error) {
	return c.do(req)
}


func (c *Client) do(req *Request) (retres *Response, reterr error) {
	if req.URL == nil {
		// req.closeBody()
		return nil, &url.Error{
			Op:  urlErrorOp(req.Method),
			Err: errors.New("http: nil Request.URL"),
		}
	}

	// 预定义一些变量
	var (
		// deadline      = c.deadline()
		reqs          []*Request
		resp          *Response
		// copyHeaders   = c.makeHeadersCopier(req)
		reqBodyClosed = false // have we closed the current req.Body?

		// Redirect behavior:
		redirectMethod string
		includeBody    bool
	)

	// 1.
	for {
		if len(reqs) > 0 {
			// For all but the first request, create the next
			// request hop and replace req.

			// 1. close resp 的 body
			// 通过检查 req 一系列的参数，在某几种场景中
			// resp.closeBody()

			// copyHeaders(req)
			return nil, nil
		}

		// 1. 将 the first req 放入 reqs
		// 2. 有可能需要 重定向之类的操作
		reqs = append(reqs, req)

		var err error
		var didTimeout func() bool


		//var shouldRedirect bool
		//redirectMethod, shouldRedirect, includeBody = redirectBehavior(req.Method, resp, reqs[0])
		//if !shouldRedirect {
		//	return resp, nil
		//}

		deadline := time.Now()
		// 这是全局唯一一次用 send 的地方
		if resp, didTimeout, err = c.send(req, deadline); err != nil {
			// c.send() always closes req.Body
			reqBodyClosed = true
			//if !deadline.IsZero() && didTimeout() {
			//	err = &httpError{
			//		// TODO: early in cycle: s/Client.Timeout exceeded/timeout or context cancellation/
			//		err:     err.Error() + " (Client.Timeout exceeded while awaiting headers)",
			//		timeout: true,
			//	}
			//}
			return nil, nil
		}

		req.closeBody()
		return nil,nil
	}
	return nil, nil
}

func (c *Client) send(req *Request, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
	if c.Jar != nil {
		for _, cookie := range c.Jar.Cookies(req.URL) {
			_ = cookie
			// req.AddCookie(cookie)
		}
	}


	//if c.Jar != nil {
	//	if rc := resp.Cookies(); len(rc) > 0 {
	//		c.Jar.SetCookies(req.URL, rc)
	//	}
	//}
	return nil, nil, nil
}

// send issues an HTTP request.
// Caller should close resp.Body when done reading from it.
func send(ireq *Request, rt RoundTripper, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
	// 1.
	// rt == nil, req.URL == nil, req.RequestURI != ""
	//


	// 2.
	setRequestCancel(ireq, rt, deadline)

	resp, err = rt.RoundTrip(ireq)
	if err != nil {

	}

	return nil, nil, nil
}
// 这个函数很关键，是 client cancel 的监听器之一。
func setRequestCancel(req *Request, rt RoundTripper, deadline time.Time) (stopTimer func(), didTimeout func() bool) {
	return nil, nil
}