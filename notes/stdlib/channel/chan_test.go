package channel

import (
	"testing"
)

func TestChan(t *testing.T) {
	// 没有缓冲的 channel-antomay
	done  := make(chan struct{})

	// 如果没有这一行就死锁了
	// close(done)
	_, ok := <-done
	t.Log(ok)
}
