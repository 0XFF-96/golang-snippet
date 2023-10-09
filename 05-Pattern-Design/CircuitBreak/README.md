### 熔断器
1. https://www.cnblogs.com/ailumiyana/p/12397372.html（学习一下笔记的写法）
2. 


### Example 
1. 参考 Hystrix 实现一个自己的滑动窗口计数器
2. 参考资源: https://github.com/afex/hystrix-go

### Backoff 退避策略
```
package grpc

import (
	"math/rand"
	"time"
)

// DefaultBackoffConfig uses values specified for backoff in
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
var (
	DefaultBackoffConfig = BackoffConfig{
		MaxDelay:  120 * time.Second,
		baseDelay: 1.0 * time.Second,
		factor:    1.6,
		jitter:    0.2,
	}
)

// backoffStrategy defines the methodology for backing off after a grpc
// connection failure.
//
// This is unexported until the gRPC project decides whether or not to allow
// alternative backoff strategies. Once a decision is made, this type and its
// method may be exported.
type backoffStrategy interface {
	// backoff returns the amount of time to wait before the next retry given
	// the number of consecutive failures.
	backoff(retries int) time.Duration
}

// BackoffConfig defines the parameters for the default gRPC backoff strategy.
type BackoffConfig struct {
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration

	// TODO(stevvooe): The following fields are not exported, as allowing
	// changes would violate the current gRPC specification for backoff. If
	// gRPC decides to allow more interesting backoff strategies, these fields
	// may be opened up in the future.

	// baseDelay is the amount of time to wait before retrying after the first
	// failure.
	baseDelay time.Duration

	// factor is applied to the backoff after each retry.
	factor float64

	// jitter provides a range to randomize backoff delays.
	jitter float64
}

func setDefaults(bc *BackoffConfig) {
	md := bc.MaxDelay
	*bc = DefaultBackoffConfig

	if md > 0 {
		bc.MaxDelay = md
	}
}

func (bc BackoffConfig) backoff(retries int) (t time.Duration) {
	if retries == 0 {
		return bc.baseDelay
	}
	backoff, max := float64(bc.baseDelay), float64(bc.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= bc.factor
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + bc.jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
```

### circuit breaker 的相关用法资料
- https://github.com/Netflix/Hystrix/wiki/How-it-Works#CircuitBreaker
- 