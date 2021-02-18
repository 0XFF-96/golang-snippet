package grpc_go

import (
	"context"
	"github.com/grpc/grpc-go/balancer"
	"github.com/liangzhiyang/annotate-grpc-go/transport"
)

func newPickerWrapper() *pickerWrapper {
	bp := &pickerWrapper{blockingCh: make(chan struct{})}
	return bp
}

func (bp *pickerWrapper) updateConnectionError(err error) {
	return
}

func (bp *pickerWrapper) connectionError() error {
	return nil
}

// updatePicker is called by UpdateBalancerState.
// It unblocks all blocked pick.
func (bp *pickerWrapper) updatePicker(p balancer.Picker) {
	return
}

//func doneChannelzWrapper(acw *acBalancerWrapper,
// done func(balancer.DoneInfo)) func(balancer.DoneInfo) {
//	return nil, nil, nil
//}

// pick returns the transport that will be used for the RPC.
// It may block in the following cases:
// - there's no picker
// - the current picker returns ErrNoSubConnAvailable
// - the current picker returns other errors and failfast is false.
// - the subConn returned by the current picker is not READY
// When one of these situations happens, pick blocks until the picker gets updated.
func (bp *pickerWrapper) pick(
	ctx context.Context,
	failfast bool,
	opts balancer.PickOptions,
	) (
	transport.ClientTransport,
	func(balancer.DoneInfo),
	error,
	) {
		// If ok == false, ac.state is not READY.
		// A valid picker always returns READY subConn. This means the state of ac
		// just changed, and picker will be updated shortly.
		// continue back to the beginning of the for loop to repick.
	return nil, nil, nil
}

func (bp *pickerWrapper) close() {
}
