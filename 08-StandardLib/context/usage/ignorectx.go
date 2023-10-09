package usage

import (
	"context"
)

type ignoreCancelContext struct {
	context.Context
}

// 忽略 Done 事件
func (icc ignoreCancelContext) Done() <-chan struct{} {
	return nil
}