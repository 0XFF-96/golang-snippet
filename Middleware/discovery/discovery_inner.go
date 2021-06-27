package discovery

import (
	"context"
)

var (
	_fetchAllURL = "http://%s/discovery/fetch/all"
)

// Protected return if service in init protect mode.
// if service in init protect mode,only support write,
// read operator isn't supported.
func (d *Discovery) Protected() bool {
	return d.protected
}

//
// syncUp populates the registry information from a peer eureka node.
// 从 eureka 中抄回来的: https://www.jianshu.com/p/f3040343e037
func (d *Discovery) syncUp() {
}

func (d *Discovery) regSelf() context.CancelFunc {
	return nil
}

func (d *Discovery) nodesproc() {
}
