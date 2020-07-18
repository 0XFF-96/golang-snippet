package example2

import (
	"context"
	"fmt"
	"time"
)

// 逆向思维一下，
// 如果我不这个函数，不想它收到超时的影响该怎么办？
// 图解 A , B, C, D, E  函数
// 当 A 函数被 cancel 了, 希望全部正在执行的函数，也一并 cancel

// 当 A 函数被 cancel 了，希望 B, C 函数也一起被 cancel, 但不希望 D, E 函数被 cancel
// gRPC: context have been cancel ! 原理？
// This example passes a context with an arbitrary deadline to tell a blocking
// function that it should abandon its work as soon as it gets to it.
func ExampleWithDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// Output:
	// context deadline exceeded
}
