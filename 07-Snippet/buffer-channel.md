### Watch Ch  设计

from: config-center 
 valueChs 的作用，用于广播信号。
```
func (d *dataContainer) WatchCh() <-chan interface{} {
	ch := make(chan interface{}, 1)
	d.mu.Lock()
	defer d.mu.Unlock()
	d.valueChs = append(d.valueChs, ch)
	return ch
}
```

### notification channel 的设计

- 这个 5 个 channel 的处理有什么问题？
- 是不是可以有更好的解决办法？将 channel 包一层？queue 处理？
- 怎么测试是否存在这种情况呢？

```
	TimingWheel struct {
		interval      time.Duration
		ticker        timex.Ticker
		slots         []*list.List
		timers        *SafeMap
		tickedPos     int
		numSlots      int
		execute       Execute

///  
		setChannel    chan timingEntry
		moveChannel   chan baseEntry
		removeChannel chan interface{}
		drainChannel  chan func(key, value interface{})
		stopChannel   chan lang.PlaceholderType
/// 
	}
```

```
func (tw *TimingWheel) run() {
	for {
		select {
		case <-tw.ticker.Chan():
			tw.onTick()
		case task := <-tw.setChannel:
			tw.setTask(&task)
		case key := <-tw.removeChannel:
			tw.removeTask(key)
		case task := <-tw.moveChannel:
			tw.moveTask(task)
		case fn := <-tw.drainChannel:
			tw.drainAll(fn)
		case <-tw.stopChannel:
			tw.ticker.Stop()
			return
		}
	}
}
```

更好的做法是：
```
// PubSub is a collection of topics.
type PubSub struct {
	cmdChan  chan cmd
	capacity int
}

type cmd struct {
	op     operation
	topics []string
	ch     chan interface{}
	msg    interface{}
}
```


### 如何实现等待队列？
- 并发 goroutine 控制
- ch 控制着 pool 中连接的数量，当取走一个时，需要 <-ch，当还回一个时，
需要 ch <- struct{}{}。
另外，还要考虑到某些失败的情况，是否需要将配额还回 ch

- redigo: Get() 连接池，是如何实现一下功能的？
	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
- 利用  buffer chan 的一些特定，简化设计。

```
		if wait {
			start = time.Now()
		}
		if ctx == nil {
			<-p.ch
		} else {
			select {
			case <-p.ch:
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		if wait {}
		
```

