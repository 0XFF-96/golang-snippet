package go_cache

import (
	"sync"
	"time"
)

type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	
	// 大锁🔒
	mu                sync.RWMutex
	onEvicted         func(string, interface{})

	// 警卫
	janitor           *janitor
}

type Item struct {
	Object     interface{}
	Expiration int64
}

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

type unexportedShardedCache struct {
	*shardedCache
}

type shardedCache struct {
	seed    uint32
	m       uint32
	cs      []*cache
	janitor *shardedJanitor
}

type shardedJanitor struct {
	Interval time.Duration
	stop     chan bool
}