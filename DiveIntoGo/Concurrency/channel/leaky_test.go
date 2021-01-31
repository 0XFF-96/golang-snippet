package channel


// 计算 goroutine 的相关数量

import (
	"fmt"
	"runtime"
	"testing"
)

func squaresNum(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

// Go scheduler also schedule goroutines on fmt.Println statement
// due to blocking I/O operation,
// however, this operation is not always blocking

func TestPrint(t *testing.T) {
	fmt.Println("main() started")
	c := make(chan int, 3)
	go squaresNum(c)

	//
	// 这行代码的重要性？
	// time.Sleep(time.Second)

	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here

	// time.Sleep(time.Second)
	fmt.Println("active goroutines", runtime.NumGoroutine())

	go squaresNum(c)

	// time.Sleep(time.Second)
	fmt.Println("active goroutines", runtime.NumGoroutine())

	c <- 5
	c <- 6
	c <- 7
	c <- 8 // blocks here

	// time.Sleep(time.Second)
	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")
}

