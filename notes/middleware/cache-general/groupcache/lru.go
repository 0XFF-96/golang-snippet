package groupcache



// 这里有并发读写问题
// 极端情况下 if c.cache == nil , 同时有两个执行的时候， 会导致覆盖。
// 丢数据，但是几率太低了。
// Add adds a value to the cache.
//func (c *Cache) Add(key Key, value interface{}) {
//	if c.cache == nil {
//		c.cache = make(map[interface{}]*list.Element)
//		c.ll = list.New()
//	}
//	if ee, ok := c.cache[key]; ok {
//		c.ll.MoveToFront(ee)
//		ee.Value.(*entry).value = value
//		return
//	}
//	ele := c.ll.PushFront(&entry{key, value})
//	c.cache[key] = ele
//	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
//		c.RemoveOldest()
//	}
//}

