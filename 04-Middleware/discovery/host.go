package discovery

import (
	"sync"
)

type hosts struct {
	hclock sync.RWMutex
	hosts  map[string]*conn // host name to conn
}
