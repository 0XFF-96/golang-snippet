package cache

import (
	"sync"
	"time"

	"github.com/Alex1996a/go-zero/core/syncx"
)

var emptyLruCache = emptyLru{}

type (
	CacheOption func(cache *Cache)

	Cache struct {
		name           string
		lock           sync.Mutex
		data           map[string]interface{}
		expire         time.Duration
		timingWheel    *TimingWheel
		lruCache       lru

		// 这个 barrier 的设计，优缺点？
		barrier        syncx.SharedCalls

		// unstable
		unstableExpiry Unstable
		stats          *cacheStat
	}

	// NOTE: 可以优化的空间
	// 1. 读写锁🔒  RWmutext
	// 2. 分 shard
	// 3. 可能( ring buffer ?)
	// 4. 元信息和存储分离
)

func NewCache(expire time.Duration, opts ...CacheOption) (*Cache, error) {
	cache := &Cache{
		data:           make(map[string]interface{}),
		expire:         expire,
		lruCache:       emptyLruCache,
		barrier:        syncx.NewSharedCalls(),
		unstableExpiry: NewUnstable(expiryDeviation),
	}

	for _, opt := range opts {
		opt(cache)
	}

	if len(cache.name) == 0 {
		cache.name = defaultCacheName
	}
	cache.stats = newCacheStat(cache.name, cache.size)

	timingWheel, err := NewTimingWheel(time.Second, slots, func(k, v interface{}) {
		key, ok := k.(string)
		if !ok {
			return
		}

		cache.Del(key)
	})
	if err != nil {
		return nil, err
	}

	cache.timingWheel = timingWheel
	return cache, nil
}

func (c *Cache) onEvict(key string) {
	// already locked
	delete(c.data, key)
	c.timingWheel.RemoveTimer(key)
}

func (c *Cache) size() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.data)
}

func (c *Cache) Del(key string) {
	c.lock.Lock()
	delete(c.data, key)
	c.lruCache.remove(key)
	c.lock.Unlock()
	c.timingWheel.RemoveTimer(key)
}

func (c *Cache) Set(key string, value interface{}) {
	c.lock.Lock()
	_, ok := c.data[key]
	c.data[key] = value
	c.lruCache.add(key)
	c.lock.Unlock()

	expiry := c.unstableExpiry.AroundDuration(c.expire)
	if ok {
		c.timingWheel.MoveTimer(key, expiry)
	} else {
		c.timingWheel.SetTimer(key, value, expiry)
	}
}