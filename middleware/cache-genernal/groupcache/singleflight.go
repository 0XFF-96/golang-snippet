package groupcache

import (
	"sync"
)

// call is an in-flight or completed Do call
type call struct {
	wg  sync.WaitGroup

	// 两个返回值
	val interface{}
	err error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}


