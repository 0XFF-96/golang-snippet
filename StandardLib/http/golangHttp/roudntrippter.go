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
	return nil, nil
}

// 底层是由 presistConn 的 roundTrip 真正实现的