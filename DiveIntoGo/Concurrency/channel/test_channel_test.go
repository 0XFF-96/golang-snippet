package channel

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

// buffered channel 的阻塞情况
// Technically, that means goroutine
// that reads buffer channel will not block until the buffer is empty
// Technically, that means goroutine that reads buffer channel will not block until the buffer is empty
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
	// c <- 4

	fmt.Println("main() stopped")
}


func TestCapacity(t *testing.T) {
	// length and capacity of a channel
	c := make(chan int, 3)
	c <- 1
	c <- 2

	fmt.Printf("Length of channel c is %v and" +
		"capacity of channel c is %v", len(c), cap(c))
	fmt.Println()
}
