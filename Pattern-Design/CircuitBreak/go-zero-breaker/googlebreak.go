package go_zero_breaker

import (
	"time"

	"github.com/Alex1996a/go-zero/core/collection"
	"github.com/Alex1996a/go-zero/core/mathx"
)

const (
	// 250ms for bucket duration
	window     = time.Second * 10
	buckets    = 40
	k          = 1.5
	protection = 5
)

// 真正的实现
// googleBreaker is a netflixBreaker pattern from google.
// see Client-Side Throttling section in https://landing.google.com/sre/sre-book/chapters/handling-overload/
type googleBreaker struct {
	k     float64
	stat  *collection.RollingWindow
	proba *mathx.Proba
}


// 优点
// 1. 滑动窗口的作用
// 2.

// 通过 	Acceptable func(err error) bool 函数，自定义哪些错误必须被计算在 breaker 的 count 里面。
// b.markSuccess()， b.markFailure() 底层利用 rolling window 的 bucket 计算
// bucket 中暴露 reduce 接口，能够自定义自己的计算函数。

// 问题：每次都需要锁🔒，会不会有冲突？
// 是不是应该【每个函数】一个单独一个锁？
func newGoogleBreaker() *googleBreaker {
	bucketDuration := time.Duration(int64(window) / int64(buckets))

	// 两个参数的作用，至关重要
	st := collection.NewRollingWindow(buckets, bucketDuration)
	return &googleBreaker{
		stat:  st,
		k:     k,
		proba: mathx.NewProba(),
	}
}


