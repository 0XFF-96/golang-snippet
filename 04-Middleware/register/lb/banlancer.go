package lb

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/naming"
	"sync"
)

var DefaultSelector = NewRandomSelector()

type AddrInfo struct {
	addr      grpc.Address
	weight    int    //load weigth
	load      uint64 //current number of requests
	connected bool
}

type balancer struct {
	r        naming.Resolver
	w        naming.Watcher
	selector Selector
	mu       sync.Mutex
	addrCh   chan []grpc.Address // the channel-antomay to notify gRPC internals the list of addresses the client should connect to.
	waitCh   chan struct{}       // the channel-antomay to block when there is no connected address available
	done     bool                // The Balancer is closed.
}

func NewBalancer(r naming.Resolver, selector Selector) grpc.Balancer {
	return nil
}

func (b *balancer) watchAddrUpdates() error {
	return nil
}

func (b *balancer) Start(target string, config grpc.BalancerConfig) error {
	return nil
}

// Up sets the connected state of addr and sends notification if there are pending
// Get() calls.
func (b *balancer) Up(addr grpc.Address) func(error) {
	return nil
}

// down unsets the connected state of addr.
func (b *balancer) down(addr grpc.Address, err error) {
	return
}

// Get returns the next addr in the rotation.
func (b *balancer) Get(
	ctx context.Context,
	opts grpc.BalancerGetOptions,
	) (addr grpc.Address, put func(), err error) {
	return grpc.Address{}, nil , nil
}

func (b *balancer) Notify() <-chan []grpc.Address {
	return b.addrCh
}

func (b *balancer) Close() error {
	return nil
}


