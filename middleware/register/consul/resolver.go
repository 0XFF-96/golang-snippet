package consul

import (
	"github.com/grpc/grpc-go/naming"
	"github.com/pkg/errors"
)

// Resolver is the implementaion of grpc.naming.Resolver
type Resolver struct {
	serviceName string
	consulAddr  string
}

// NewResolver return ConsulResolver with service name
func NewResolver(serviceName string, consulAddr string) *Resolver {
	return &Resolver{serviceName: serviceName, consulAddr: consulAddr}
}

// Resolve to resolve the service from consul
func (c *Resolver) Resolve(target string) (naming.Watcher, error) {
	if c.serviceName == "" {
		return nil, errors.New("no service name provided")
	}
	watcher := newWatcher(c.serviceName, c.consulAddr)
	return watcher, nil
}


