package chapter6

import (
	"fmt"
	"testing"
)

// 为什么需要 Work Stealing
// 用 p1, p2 和 task1, task2, task3, task4
//
//


// 第一个改进
// load-balancing problems that
// maybe a FIFO queue can help with.
// but  cenralized queue
// 2、cache locality and processor utilization

// let's go through the rules of how a
// work-stealing algorithm operates with distributed
// queues
// fork-join model for concurrency
func TestFun(t *testing.T){
	var fib func(n int) <-chan int
	fib = func(n int) <- chan int {
		result := make(chan int)
		go func(){
			defer close(result)
			if n <=2 {
				result <-1
				return
			}
			result <- <- fib((n-1)) + <-fib(n-2)
		}()
		return result
	}

	fmt.Printf("fib(4) = %d", <-fib(4))
}
// T1 call stack | T1 work queue | T2 call stack | T2 work queue
// ...
// ...
// ...
// distributed-queue work-stealing algorithm
//

// tasks stealing
// continuation stealing

