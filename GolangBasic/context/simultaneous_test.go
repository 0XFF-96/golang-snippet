package context

import (
	"context"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 这个例子
// 针对如何对并发函数进行测试，有很大的启发作用
//
func TestSimultaneous(t *testing.T) {
	root, cancel := context.WithCancel(context.Background())
	m := map[context.Context]context.CancelFunc{root: cancel}
	q := []context.Context{root}
	// Create a tree of contexts.
	for len(q) != 0 && len(m) < 100 {
		parent := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			ctx, cancel := context.WithCancel(parent)
			m[ctx] = cancel
			q = append(q, ctx)
		}
	}
	// Start all the cancels in a random order.
	var wg sync.WaitGroup
	wg.Add(len(m))
	for _, cancel := range m {
		go func(cancel context.CancelFunc) {
			cancel()
			wg.Done()
		}(cancel)
	}
	// Wait on all the contexts in a random order.
	for ctx := range m {
		select {
		case <-ctx.Done():
		case <-time.After(1 * time.Second):
			buf := make([]byte, 10<<10)
			n := runtime.Stack(buf, true)
			t.Fatalf("timed out waiting for <-ctx.Done(); stacks:\n%s", buf[:n])
		}
	}
	// Wait for all the cancel functions to return.
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		buf := make([]byte, 10<<10)
		n := runtime.Stack(buf, true)
		t.Fatalf("timed out waiting for cancel functions; stacks:\n%s", buf[:n])
	}
}

// 这里是测试，
// 发生死锁的情况下，
// 能不能在 1 s 内恢复
func TestInterlockedCancels(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)

	// ---- 可能导致死锁的代码------------
	go func() {
		parent.Done()
		cancelChild()
	}()
	cancelParent()
	// ---------------------------------

	select {
	case <-child.Done():
	case <-time.After(1 * time.Second):
		buf := make([]byte, 10<<10)
		n := runtime.Stack(buf, true)
		t.Fatalf("timed out waiting for child.Done(); stacks:\n%s", buf[:n])
	}
}

// TODO:@Me
// 后续把这部分代码，移动到专门学习 Channel 的地方去。
// 参考文章: https://colobu.com/2016/04/14/Golang-Channels/
func TestCloseCh(t *testing.T) {
	// 能够从一个 close 的 chan 读取到值
	ch := make(chan int)
	close(ch)
	t.Log(<-ch)

	// 判断一个 chan 是否被关闭
	v, ok := <-ch
	t.Log(v)
	t.Log(ok)

	// panic: 往一个已经关闭到 channel 写东西
	// panic: send on closed channel [recovered]
	// ch <- 1
}
