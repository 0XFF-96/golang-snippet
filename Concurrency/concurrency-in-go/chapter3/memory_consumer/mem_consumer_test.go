package memory_consumer

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

)

func TestMemConsumed(t *testing.T){
	memStat := func()uint64{
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func(){wg.Done(); <-c}

	const numGoruntines = 1e5
	wg.Add(numGoruntines)
	before := memStat()
	for i:=numGoruntines;i > 0;i--{
		go noop()
	}
	wg.Wait()
	after := memStat()
	fmt.Printf("%.3fkb", float64(after - before)/numGoruntines/1000)
}

