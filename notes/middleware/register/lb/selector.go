package lb

import (
	"context"

	"google.golang.org/grpc"
)

type Selector interface {
	Add(addr grpc.Address) error
	Delete(addr grpc.Address) error
	Up(addr grpc.Address) (cnt int, connected bool)
	Down(addr grpc.Address) error
	AddrList() []grpc.Address
	Get(ctx context.Context) (grpc.Address, error)
	Put(addr string) error
}


