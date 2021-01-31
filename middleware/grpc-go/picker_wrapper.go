package grpc_go

import (
	"sync"

	"github.com/grpc/grpc-go/balancer"
)

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
// actions and unblock when there's a picker update.
type pickerWrapper struct {
	mu         sync.Mutex
	done       bool
	blockingCh chan struct{}
	picker     balancer.Picker

	// The latest connection happened.
	connErrMu sync.Mutex
	connErr   error
}






