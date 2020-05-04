package example

import (
	"context"
	"fmt"
)

// 来自于官方文档
// 用法1: 用于防止 goroutine leak 泄露
// 哪一行代码，会导致 goroutine 泄露？为什么？
// 画出下面代码的程序时序图？
// 下面代码的完整生命周期
// 知识点:1 fork-join 模型
// dst 是沟通主携程和 gen 携程的桥梁。（抽象一下）
// This example demonstrates the use of a cancelable context to prevent a
// goroutine leak. By the end of the example function, the goroutine started
// by gen will return without leaking.
func ExampleWithCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
					// 能否用图表的形式讲解一下，
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
