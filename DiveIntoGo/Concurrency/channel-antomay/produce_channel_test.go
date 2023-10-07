package channel_antomay

import (
	"fmt"
	"testing"
)

func produceSquares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel-antomay
}

func TestReceiveChan(t *testing.T) {
	fmt.Println("main() started")
	c := make(chan int)

	go produceSquares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for val := range c {
		fmt.Println(val)
	}

	// time.After()
	fmt.Println("main() stopped")
}
