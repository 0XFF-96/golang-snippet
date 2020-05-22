package once

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T){
	var count int
	increment := func(){
		count ++
	}

	var incr sync.WaitGroup
	var once sync.Once
	incr.Add(100)
	for i:=0;i<100;i++{
		go func(){
			once.Do(increment)
		}()
	}
	fmt.Println(count)
}
// 查看 sync.Once 被使用了多少次的技巧...
// grep -ir sync.Once $(go env GRoot)/src | wc -l

func TestOnceV2(t *testing.T){
	var count int
	increment := func(){ count++}
	decrement := func(){ count--}

	var once sync.Once
	//  /*-------------------*/
	// once only count the number of do
	// 只会执行最开始的 once.Do
	// 如何 once.Do 内部是如何计数的？
	// /*-------------------*/
	once.Do(decrement)
	once.Do(increment)
	once.Do(increment)
	fmt.Println(count)
}

// 这个程序发生死🔒了
// 这个程序会发送什么效果？
// 为什么会发送 deadlock ?
// 他们互相持了什么资源？
// in this case a circular reference.
func TestOnceV3(t *testing.T){
	var onceA, onceB sync.Once
	var initB func()

	initA := func(){onceB.Do(initB)}
	initB = func(){onceA.Do(initA)}

	onceA.Do(initA)
}
