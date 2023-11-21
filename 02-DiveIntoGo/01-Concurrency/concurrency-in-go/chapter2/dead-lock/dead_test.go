package dead_lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// review: 一次

// page24
// https://github.com/golang/go/issues/13759.
// detect runtime deadlocks
type value struct {
	mu sync.Mutex
	value int
}


// Question: 请画出下面程序的执行流程图
// 出现了什么问题？
// 在哪几行代码出现的？
// 什么是 race condition ?
func TestDeadlock(t *testing.T){
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value){
		defer  wg.Done()
		v1.mu.Lock()
		defer  v1.mu.Unlock()
		time.Sleep(2*time.Second) // 这行才有可能触发 deadlock
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value + v2.value)
	}
	var a, b value
	wg.Add(2)
	// 规律 a,b
	// b, a
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
