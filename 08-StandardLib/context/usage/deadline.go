package usage

import (
	"context"
	"time"
)

// From Go-zero
func ShrinkDeadline(ctx context.Context, timeout time.Duration) (context.Context, func()) {
	if deadline, ok := ctx.Deadline(); ok {
		leftTime := time.Until(deadline)
		if leftTime < timeout {
			timeout = leftTime
		}
	}

	return context.WithDeadline(ctx, time.Now().Add(timeout))
}
