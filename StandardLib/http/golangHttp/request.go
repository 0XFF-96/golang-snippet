package golangHttp

import (
	"net/url"
)

type Request struct {
	// 1.
	URL *url.URL

	// 2.
	Method string
}


func (r *Request) closeBody() {
}