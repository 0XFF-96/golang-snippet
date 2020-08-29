package _map

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type Map struct {
	mu sync.Mutex

	read atomic.Value // readOnly
	dirty map[interface{}]*entry
	misses int
}

// An entry is a slot in the map corresponding to a particular key.
type entry struct {
	// entry 有三种情况
	// If p == nil
	// If p == expunged
	// the entry is valid and recorded in m.read.m[key] and, if m.dirty
	// != nil, in m.dirty[key].
	// An entry's associated value can be updated by atomic replacement
	p unsafe.Pointer // *interface{}
}
