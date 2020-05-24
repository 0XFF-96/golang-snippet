package singleton

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		// 并不是并发安全的代码
		// 因为在 其他 goroutine 有可能同时调用
		instance = &singleton{} // <--- NOT THREAD SAFE
	}

	// The reason this is bad is that if references to the singleton instance
	// are being held around through the code,
	// there could be potentially multiple instances of the type with different states
	return instance
}

// The Aggressive Locking

var mu sync.Mutex

func GetInstanceV2() *singleton {
	mu.Lock() // <--- Unnecessary locking if instance already created
	defer mu.Unlock()
	// Sync.Mutex and acquiring the Lock before creating the singleton instance
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// Check-Lock-Check Pattern
// if check {
// lock() {
//    if check(){
//      // your code
// }
//
// 为什么需要第二次 check 的原因
//  But butween the first check and the acquisition of the exclusive lock
//  there could have been another thread that did acquire the lock,
//  therefore we would need to check again inside the lock to
//  avoid replacing the instance with another one

// 更好的一个方法
func GetInstanceV3() *singleton {
	if instance == nil { // <-- Not yet perfect. since it's not fully atomic
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

// 因为 compiler optimizations there is not an atomic check on the instance store state
//
// 改进方法: compiler optimizations there is not an atomic check on the instance store state
//

var initialized uint32

func GetInstanceV4() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// An Idiomatic Singleton Approach in Go
//

var once sync.Once

func GetInstanceV5() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
