package leakybucket

import (
	"sync"
	"time"
)

// RateLimiter 是一个过滤器，
// 用于检查 itemCsot 是否在 rate limits
// 的限制下
type RateLimiter interface {
	CheckCredit(itemCost float64) bool
}


type rateLimiter struct {
	sync.Mutex

	creditsPerSecond float64
	balance float64
	maxBalance float64

	lastTick time.Time

	timeNow func() time.Time
}

// NewRateLimiter基于漏斗算法创建了一个新的速率限制器，公式为
// 每次调用 CheckCredit（）方法时都会按比例成比例的金额补充的贷方余额
// 自上次滴答后经过的时间，最多不超过creditsPerSecond。调用CheckCredit（）会产生成本
// 我们要用余额支付的项目。如果余额超过了物料成本，则该物料为“已购买”
// 余额减少，以true的返回值表示。否则，余额保持不变并返回false。
//
// 可以通过实例化 Rate Limiter 来限制服务发出的消息的速率。
// 每秒允许服务发出的最大消息数，并为每个消息调用 CheckCredit（1.0）
// 确定消息是否在速率限制内。
//
// 通过将 creditsPerSecond 设置为所需的吞吐量，它还可用于限制流量的字节数
// 以字节/秒为单位，并使用实际消息大小调用CheckCredit（）。


func NewRateLimiter(creditsPerSecond, maxBalance float64) RateLimiter {
	return &rateLimiter{
		creditsPerSecond: creditsPerSecond,
		balance:          maxBalance,
		maxBalance:       maxBalance,
		lastTick:         time.Now(),
		timeNow:          time.Now}
}

func (b *rateLimiter) CheckCredit(itemCost float64) bool {
	b.Lock()
	defer b.Unlock()
	// calculate how much time passed since the last tick, and update current tick

	// 上次调用到现在的间隔，
	currentTime := b.timeNow()
	elapsedTime := currentTime.Sub(b.lastTick)
	b.lastTick = currentTime

	// calculate how much credit have we accumulated since the last tick

	// 我们从上次调用到现在到时间中累积了多少 credit
	b.balance += elapsedTime.Seconds() * b.creditsPerSecond
	if b.balance > b.maxBalance {
		b.balance = b.maxBalance
	}
	// if we have enough credits to pay for current item, then reduce balance and allow
	// 我们有足够到 credit
	if b.balance >= itemCost {
		b.balance -= itemCost
		return true
	}
	return false
}