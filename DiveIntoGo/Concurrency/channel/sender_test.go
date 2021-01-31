package channel

import (
	"fmt"
	"testing"
)

func sender(c chan int) {
	c <- 1 // len 1, cap 3
	c <- 2 // len 2, cap 3
	c <- 3 // len 3, cap 3
	c <- 4 // <- goroutine blocks here
	close(c)
}

func TestSendBlock(t *testing.T) {
	c := make(chan int, 3)

	go sender(c)

	fmt.Printf("Length of channel c is %v and " +
		"capacity of channel c is %v\n", len(c), cap(c))

	// read values from c (blocked here)
	for val := range c {
		fmt.Printf("Length of channel c after " +
			"value '%v' read is %v\n", val, len(c))
	}
}

