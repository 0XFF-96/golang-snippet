package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// otherContext is a Context that's not one of the types defined in context.go.
// This lets us test code paths that differ based on the underlying type of the
// Context.
type otherContext struct {
	context.Context
}

// 理解这段代码
// 对于 golang 的 context 是如何被 cancel 会有很大的帮助！
func TestContextCancel(t *testing.T) {
	c1, cancel := context.WithCancel(context.Background())

	if got, want := fmt.Sprint(c1), "context.Background.WithCancel"; got != want {
		t.Errorf("c1.String() = %q want %q", got, want)
	}

	o := otherContext{c1}
	c2, _ := context.WithCancel(o)
	contexts := []context.Context{c1, o, c2}

	for i, c := range contexts {
		if d := c.Done(); d == nil {
			t.Errorf("c[%d].Done() == %v want non-nil", i, d)
		}
		if e := c.Err(); e != nil {
			t.Errorf("c[%d].Err() == %v want nil", i, e)
		}

		select {
		case x := <-c.Done():
			t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
		default:
		}
	}

	cancel()
	time.Sleep(100 * time.Millisecond) // let cancellation propagate

	for i, c := range contexts {
		select {
		case <-c.Done():
		default:
			t.Errorf("<-c[%d].Done() blocked, but shouldn't have", i)
		}
		if e := c.Err(); e != context.Canceled {
			t.Errorf("c[%d].Err() == %v want %v", i, e, context.Canceled)
		}
	}
}

func TestParentFinishesChild(t *testing.T) {
	// Context tree:
	// parent -> cancelChild
	// parent -> valueChild -> timerChild
	parent, cancel := context.WithCancel(context.Background())
	cancelChild, stop := context.WithCancel(parent)
	defer stop()
	valueChild := context.WithValue(parent, "key", "value")
	timerChild, stop := context.WithTimeout(valueChild, 10000*time.Hour)
	defer stop()

	select {
	case x := <-parent.Done():
		t.Errorf("<-parent.Done() == %v want nothing (it should block)", x)
	case x := <-cancelChild.Done():
		t.Errorf("<-cancelChild.Done() == %v want nothing (it should block)", x)
	case x := <-timerChild.Done():
		t.Errorf("<-timerChild.Done() == %v want nothing (it should block)", x)
	case x := <-valueChild.Done():
		t.Errorf("<-valueChild.Done() == %v want nothing (it should block)", x)
	default:
	}

	// 思考🤔
	// 此时，内存中有什么数据表示 parent、cancelChild、valueChild、timerChild 的数据结构？
	// 你能把它图形化出来吗？

	// 提示一下
	// A cancelCtx can be canceled. When canceled, it also cancels any children
	// that implement canceler.
	//type cancelCtx struct {
	//	Context   // -> 这个是啥？ 其实就是它的 parent context, 由哪一个 parent context 创建的
	// 			  //
	//
	//	mu       sync.Mutex            // protects following fields
	//	done     chan struct{}         // created lazily, closed by first cancel call
	//	children map[canceler]struct{} // set to nil by the first cancel call
	//	err      error                 // set to non-nil by the first cancel call
	//}
	//// The parent's children should contain the two cancelable children.
	// 因为有些数据结构不导出，所以省略...
	//
}
