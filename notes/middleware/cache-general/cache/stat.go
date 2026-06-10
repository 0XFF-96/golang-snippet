package cache

import (
	"sync/atomic"
	"time"

	"github.com/bilibili/overlord/pkg/log"
)

type cacheStat struct {
	name         string
	hit          uint64
	miss         uint64
	sizeCallback func() int
}


func newCacheStat(name string, sizeCallback func() int) *cacheStat {
	st := &cacheStat{
		name:         name,
		sizeCallback: sizeCallback,
	}
	go st.statLoop()
	return st
}

func (cs *cacheStat) IncrementHit() {
	atomic.AddUint64(&cs.hit, 1)
}

func (cs *cacheStat) IncrementMiss() {
	atomic.AddUint64(&cs.miss, 1)
}

func (cs *cacheStat) statLoop() {
	ticker := time.NewTicker(statInterval)
	defer ticker.Stop()

	for range ticker.C {
		hit := atomic.SwapUint64(&cs.hit, 0)
		miss := atomic.SwapUint64(&cs.miss, 0)
		total := hit + miss
		if total == 0 {
			continue
		}
		percent := 100 * float32(hit) / float32(total)
		// it can be replace with our own logger
		// grafana 监控起来
		log.Infof("cache(%s) - qpm: %d, hit_ratio: %.1f%%, elements: %d, hit: %d, miss: %d",
			cs.name, total, percent, cs.sizeCallback(), hit, miss)
	}
}