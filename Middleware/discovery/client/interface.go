package client

import (
	"context"
)

// Resolver resolve naming service
type Resolver interface {
	Fetch() (*InstancesInfo, bool)
	Watch() <-chan struct{}
	Close() error
}

// Registry Register an instance and renew automatically.
type Registry interface {
	Register(ins *Instance) (cancel context.CancelFunc, err error)
	Close() error
}

// Builder resolver builder.
type Builder interface {
	Build(id string) Resolver
	Scheme() string
}

