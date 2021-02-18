package storeCache

import (
	"time"

	"github.com/Alex1996a/go-zero/core/hash"
)

type (
	Cache interface {
		DelCache(keys ...string) error
		GetCache(key string, v interface{}) error
		IsNotFound(err error) bool
		SetCache(key string, v interface{}) error
		SetCacheWithExpire(key string, v interface{}, expire time.Duration) error
		Take(v interface{}, key string, query func(v interface{}) error) error
		TakeWithExpire(v interface{}, key string, query func(v interface{}, expire time.Duration) error) error
	}

	cacheCluster struct {
		dispatcher  *hash.ConsistentHash
		errNotFound error
	}
)

