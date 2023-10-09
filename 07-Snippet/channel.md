### channel usage 

- 当存在快慢程序调用时，如果进行操作取消的处理？
```
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	foo(ctx)
}

func foo(ctx context.Context) {
	ch := make(chan int, 1)
	go func() {
		slow()
		ch <- 1
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ch:
			return
		}
	}
}

func slow() {
	n := rand.Intn(3)
	fmt.Printf("sleep %ds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
```