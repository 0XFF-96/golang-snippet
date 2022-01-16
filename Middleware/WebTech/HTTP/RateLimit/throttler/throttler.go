package throttler

// Throttler is used to rate limits operations. For different-struct-memory, given how debug spans
// are always sampled, a throttler can be enabled per client to rate limit the amount
// of debug spans a client can start.
type Throttler interface {
	// IsAllowed determines whether the operation should be allowed and not be
	// throttled.
	IsAllowed(operation string) bool
}


// 用于测试, 或者其他用途的 DefaultThrottler
// DefaultThrottler doesn't throttle at all.
type DefaultThrottler struct{}

// IsAllowed implements Throttler#IsAllowed.
func (t DefaultThrottler) IsAllowed(operation string) bool {
	return true
}
