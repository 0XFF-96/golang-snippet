package go_zero_hash

import (
	"sync"

	"github.com/spaolacci/murmur3"
)

var Placeholder PlaceholderType


func Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}


type (
	GenericType     = interface{}
	PlaceholderType = struct{}
)

const (
	TopWeight = 100

	minReplicas = 100
	prime       = 16777619
)

// 虚拟节点的概念？
type (
	HashFunc func(data []byte) uint64

	ConsistentHash struct {
		hashFunc HashFunc
		replicas int
		keys     []uint64
		ring     map[uint64][]interface{}
		nodes    map[string]PlaceholderType
		lock     sync.RWMutex
	}
)

func NewConsistentHash() *ConsistentHash {
	return NewCustomConsistentHash(minReplicas, Hash)
}

func NewCustomConsistentHash(replicas int, fn HashFunc) *ConsistentHash {
	if replicas < minReplicas {
		replicas = minReplicas
	}

	if fn == nil {
		fn = Hash
	}

	return &ConsistentHash{
		hashFunc: fn,
		replicas: replicas,
		ring:     make(map[uint64][]interface{}),
		nodes:    make(map[string]PlaceholderType),
	}
}
