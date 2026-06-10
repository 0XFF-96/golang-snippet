package golangHttp

import (
	"net/url"
	"sync"
)

// connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.


// A connect method may be of the following types:
//
//	connectMethod.key().String()      Description
//	------------------------------    -------------------------
//	|http|foo.com                     http directly to server, no proxy
//	|https|foo.com                    https directly to server, no proxy
//	|https,h1|foo.com                 https directly to server w/o HTTP/2, no proxy
//	http://proxy.com|https|foo.com    http to proxy, then CONNECT to foo.com
//	http://proxy.com|http             http to proxy, http to anywhere after that
//	socks5://proxy.com|http|foo.com   socks5 to proxy, then http to foo.com
//	socks5://proxy.com|https|foo.com  socks5 to proxy, then https to foo.com
//	https://proxy.com|https|foo.com   https to proxy, then CONNECT to foo.com
//	https://proxy.com|http            https to proxy, http to anywhere after that
//
type connectMethod struct {
	proxyURL     *url.URL // nil for no proxy, else full proxy URL
	targetScheme string   // "http" or "https"
	// If proxyURL specifies an http or https proxy, and targetScheme is http (not https),
	// then targetAddr is not included in the connect method key, because the socket can
	// be reused for different targetAddr values.
	targetAddr string
	onlyH1     bool // whether to disable HTTP/2 and force HTTP/1
}

// A Header represents the key-value pairs in an HTTP header.
//
// The keys should be in canonical form, as returned by
// CanonicalHeaderKey.
type Header map[string][]string

type transportRequest struct {
	*Request                        // original request, not to be mutated
	extra    Header                 // extra headers to write, or nil
	// trace    *httptrace.ClientTrace // optional

	mu  sync.Mutex // guards err
	err error      // first setError value for mapRoundTripError to consider
}