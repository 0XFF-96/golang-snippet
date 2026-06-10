package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 你能否画出这个程序的执行流程图？
// 你能否指出这个程序的临界区在哪里？
// 这个程序可以怎么改？
// 这个程序的含义是什么？
func TestCond(t *testing.T){
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration){
		time.Sleep(delay)

		c.L.Lock()
		queue = queue[1:]
		fmt.Println("remove from Queue")
		c.L.Unlock()
		// 这句尤其重要
		c.Signal()
	}

	for i:=0; i < 10; i++{
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("add to queue")

		queue = append(queue, struct{}{})
		go removeFromQueue(1 *time.Second)
		c.L.Unlock()
	}
}


// 顺序是否固定？
// 程序流程图？
// Broadcast 的作用？
// wait 的两个作用是什么？
// 不是太明白它的解析...
func TestConBroadcast(t *testing.T){
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked:sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()){
		var goroutineRuning sync.WaitGroup
		goroutineRuning.Add(1)
		go func(){
			goroutineRuning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRuning.Wait()
	}

	var clickRegisterd sync.WaitGroup
	clickRegisterd.Add(3)

	subscribe(button. Clicked, func(){
		fmt.Println("Maximizing window.")
		clickRegisterd.Done()
	})

	subscribe(button.Clicked, func(){
		fmt.Println("Displaying annoying dialog box!")
		clickRegisterd.Done()
	})
	subscribe(button.Clicked, func(){
		fmt.Println("Mouse clicked")
		clickRegisterd.Done()
	})
	button.Clicked.Broadcast()
	clickRegisterd.Wait()
}

