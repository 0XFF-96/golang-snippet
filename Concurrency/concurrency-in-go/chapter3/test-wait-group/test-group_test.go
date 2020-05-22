package test_wait_group

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T){
	var wg sync.WaitGroup
	// add 一定要在 goroutine 之外
	// 而不能在生成的 goroutine ...
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("first go ")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("sencond Go")
		time.Sleep(2)
	}()

	wg.Wait()
}

func TestWaitGroupV2(t *testing.T) {
	// Add called to track a group of goroutines all at once.
	// I usually do this before for loops like this
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Println("id:", id)
	}
	var wg sync.WaitGroup
	const numGreeters = 5
	wg.Add(numGreeters)
	// 这个结果的顺序是固定的吗？不是固定的...
	// 为什么？
	for i := 0; i < numGreeters; i++ {
		fmt.Println("I:", i)
		go hello(&wg, i+1)
		// wg.Wait() 会 deadlock!
	}
	wg.Wait()
}