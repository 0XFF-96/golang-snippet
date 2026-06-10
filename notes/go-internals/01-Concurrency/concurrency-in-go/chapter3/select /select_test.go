package select_

import (
	"fmt"
	"testing"
	"time"
)

// 对照着 channel-antomay 的相关状态
// 从一个 close(channel-antomay) 是可以进行 read 操作的
// 只不过，是 read 到的是 defaul value，and Ok == false
func TestSelect(t *testing.T){
	c1 := make(chan interface{})
	// 这个 close 语句的作用是什么？
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	// close
	// 这里面 close 语句的作用是干什么用的
	var c1Count, c2Count int
	for i := 1000; i >=0;i--{
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++

		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func TestSelectV1(t *testing.T){
	var c1, c2 <- chan interface{}
	var c3 chan<-interface{}

	select {
	case <-c1:
	// Do something
	case <- c2:
	// Do something
	case c3<-struct{}{}:
		// Do something
	}
}

func TestSelectV2(t *testing.T){
	start := time.Now()
	c := make(chan interface{})

	go func(){
		time.Sleep(5 *time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblock %v later", time.Since(start))
	}
}

func TestSelectTimeOut(t *testing.T){
	done := make(chan interface{})

	go func(){
		time.Sleep(5 *time.Second)
		close(done)
	}()

	workCounter :=0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 *time.Second)
	}
	fmt.Printf("Achieved %v cycles of work befoer signalled to stop .\n", workCounter)
}



