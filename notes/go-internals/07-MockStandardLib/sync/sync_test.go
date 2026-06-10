package sync

import (
	"sync"
	"testing"
)

type Small struct {
	a int
}

var pool = sync.Pool{
	New: func() interface{} { return new(Small) },
}

//go:noinline
func inc(s *Small) { s.a++ }

func BenchmarkWithoutPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = &Small{ a: 1, }
			b.StopTimer(); inc(s); b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = pool.Get().(*Small)
			s.a = 1
			b.StopTimer(); inc(s); b.StartTimer()
			pool.Put(s)
		}
	}
}


// The workflow of sync pool

// 1.
//func init() {
//	runtime_registerPoolCleanup(poolCleanup)
//}


// 2.
//
// example: https://pkg.go.dev/sync#Pool


// 3.
// For each sync.Pool we create,
// go generates an internal pool poolLocal attached to each processor.

// 4. New lock-free pool and victim cache
// https://github.com/golang/go/commit/d5fd2dd6a17a816b7dfd99d4df70a85f1bf0de31#diff-491b0013c82345bf6cfa937bd78b690d

