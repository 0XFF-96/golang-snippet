package discovery

import (
	"sync"
)

// Registry handles replication of all operations to peer Discovery nodes to keep them all in sync.
type Registry struct {
	appm  map[string]*Apps // appid-env -> apps
	aLock sync.RWMutex

	conns     map[string]*hosts // region.zone.env.appid-> host
	cLock     sync.RWMutex
	scheduler *scheduler
	gd        *Guard
}


