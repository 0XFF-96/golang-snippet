package lb

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

type RandomSelector struct {
	baseSelector
	r *rand.Rand
}

func (r *RandomSelector) Add(addr grpc.Address) error {
	panic("implement me")
}

func NewRandomSelector() Selector {
	return &RandomSelector{
		r:            rand.New(rand.NewSource(time.Now().UnixNano())),
		baseSelector: baseSelector{addrMap: make(map[string]*AddrInfo)},
	}
}

func (r *RandomSelector) Get(ctx context.Context) (addr grpc.Address, err error) {
	if len(r.addrs) == 0 {
		return addr, errors.New("addr list is emtpy")
	}

	size := len(r.addrs)
	idx := r.r.Int() % size

	for i := 0; i < size; i++ {
		addr := r.addrs[(idx+i)%size]
		if addrInfo, ok := r.addrMap[addr]; ok {
			if addrInfo.connected {
				addrInfo.load++
				return addrInfo.addr, nil
			}
		}
	}
	return addr, errors.New("addr list is emtpy")
}


func (b *baseSelector) Delete(addr grpc.Address) error {
	return nil 
}

func (b *baseSelector) Up(addr grpc.Address) (cnt int, connected bool) {
	return 0, false
}

func (b *baseSelector) Down(addr grpc.Address) error {
	return nil
}

func (b *baseSelector) AddrList() []grpc.Address {
	return nil
}

func (b *baseSelector) Get(ctx context.Context) (addr grpc.Address, err error) {
	return
}

func (b *baseSelector) Put(addr string) error {
	return nil
}