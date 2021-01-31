package discovery

import (
	"sync"
	"sync/atomic"

	"go-common/library/log"
)

const (
	_percentThreshold float64 = 0.85
)

// Guard count the renew of all operations for self protection
// 自我保护
// 过载保护

// 这四个指标式是什么意思呢？
type Guard struct {
	expPerMin    int64
	expThreshold int64
	facInMin     int64
	facLastMin   int64
	lock         sync.RWMutex
}

func (g *Guard) setExp(cnt int64) {
	g.lock.Lock()
	g.expPerMin = cnt * 2
	g.expThreshold = int64(float64(g.expPerMin) * _percentThreshold)
	g.lock.Unlock()
}

func (g *Guard) incrExp() {
	g.lock.Lock()
	g.expPerMin = g.expPerMin + 2
	g.expThreshold = int64(float64(g.expPerMin) * _percentThreshold)
	g.lock.Unlock()
}

func (g *Guard) updateFac() {
	atomic.StoreInt64(&g.facLastMin, atomic.SwapInt64(&g.facInMin, 0))
}

func (g *Guard) decrExp() {
	g.lock.Lock()
	if g.expPerMin > 0 {
		g.expPerMin = g.expPerMin - 2
		g.expThreshold = int64(float64(g.expPerMin) * _percentThreshold)
	}
	g.lock.Unlock()
}

func (g *Guard) incrFac() {
	atomic.AddInt64(&g.facInMin, 1)
}

// protected 函数
// 用于保护作用
func (g *Guard) ok() (is bool) {
	is = atomic.LoadInt64(&g.facLastMin) < atomic.LoadInt64(&g.expThreshold)
	if is {
		log.Warn("discovery is protected, the factual renews(%d) less than expected renews(%d)", atomic.LoadInt64(&g.facLastMin), atomic.LoadInt64(&g.expThreshold))
	}
	return
}
