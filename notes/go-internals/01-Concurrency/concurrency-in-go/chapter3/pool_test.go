package chapter3

import (
	"fmt"
	"sync"
	"testing"
)

// 相关函数//
// 避免频繁分配内存
// 节省对象创建的开销
func TestPool(t *testing.T){
	myPool := &sync.Pool{
		New:func()interface{}{
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

func TestPoolV2(t *testing.T){
	// 试一下调整这个参数
	// 猜一下，那个数字会发生变化？
	// runtime.GOMAXPROCS(4)
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// Seed the pool with 4kb
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024*1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0;i--{
		go func(){
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	// 和 CPU 核心数有关
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}