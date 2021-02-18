package cache


func (c *Cache) Take(key string, fetch func() (interface{}, error)) (interface{}, error) {
	if val, ok := c.doGet(key); ok {
		c.stats.IncrementHit()
		return val, nil
	}

	var fresh bool

	// sharedCalls 概念
	// SharedCalls lets the concurrent calls with the same key to share the call result.
	// For example, A called F, before it's done, B called F. Then B would not execute F,
	// and shared the result returned by F which called by A.
	// The calls with the same key are dependent, concurrent calls share the returned values.
	// A ------->calls F with key<------------------->returns val
	// B --------------------->calls F with key------>returns val

	val, err := c.barrier.Do(key, func() (interface{}, error) {
		// because O(1) on map search in memory, and fetch is an IO query
		// so we do double check, cache might be taken by another call
		if val, ok := c.doGet(key); ok {
			return val, nil
		}

		v, e := fetch()
		if e != nil {
			return nil, e
		}

		fresh = true
		c.Set(key, v)
		return v, nil
	})
	if err != nil {
		return nil, err
	}

	if fresh {
		c.stats.IncrementMiss()
		return val, nil
	}

	// got the result from previous ongoing query
	c.stats.IncrementHit()
	return val, nil
}

func (c *Cache) doGet(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.data[key]
	if ok {
		c.lruCache.add(key)
	}

	return value, ok
}
