package client

import (
	"sync/atomic"
)

type appInfo struct {
	resolver map[*Resolve]struct{}
	zoneIns  atomic.Value
	lastTs   int64 // latest timestamp
}

// Resolve discveory resolver.
type Resolve struct {
	id    string
	event chan struct{}
	d     *Discovery
}

// Fetch fetch resolver instance.
func (r *Resolve) Fetch() (ins *InstancesInfo, ok bool) {
	return
}

// Close close resolver.
func (r *Resolve) Close() error {
}
