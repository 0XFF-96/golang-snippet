package channel_antomay

import (
	"fmt"
	"testing"
)

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

// buffered channel-antomay 的阻塞情况
// Technically, that means goroutine
// that reads buffer channel-antomay will not block until the buffer is empty
// Technically, that means goroutine that reads buffer channel-antomay
// will not block until the buffer is empty
func TestChannelBlocking(t *testing.T) {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3

	// 👇代码对理解 buffered/unbuffered 和 blocking 很关键
	// this code of line will blocks the main goroutine
	// and schedule to squares, squares will drain all the value
	c <- 4

	// 1. 当不加上 c <- 4 时，主函数不会被阻塞，squares 的计算结果
	// 2. 当加上了 c <-4 之后，由于 buffer channel-antomay 当原因， main() goroutine 被阻塞了
	fmt.Println("main() stopped")
}


func TestCapacity(t *testing.T) {
	// length and capacity of a channel-antomay
	c := make(chan int, 3)
	c <- 1
	c <- 2

	fmt.Printf("Length of channel-antomay c is %v and" +
		"capacity of channel-antomay c is %v", len(c), cap(c))
	fmt.Println()
}
