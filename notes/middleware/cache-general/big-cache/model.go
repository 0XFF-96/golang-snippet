package big_cache

import (
	"sync"
	"time"

	"github.com/allegro/bigcache/queue"
)

// BigCache is fast, concurrent, evicting cache created to keep big number of entries without impact on performance.
// It keeps entries on heap but omits GC for them. To achieve that, operations take place on byte arrays,
// therefore entries (de)serialization in front of the cache will be needed in most use cases.
type BigCache struct {
	shards       []*cacheShard
	lifeWindow   uint64
	// clock        clock
	hash         Hasher
	config       Config
	shardMask    uint64
	maxShardSize uint32
	close        chan struct{}
}

type cacheShard struct {
	hashmap     map[uint64]uint32

	// 重点
	// 1. 如何存储 byte data ?
	// 2. 流程图？
	entries     queue.BytesQueue

	// 有什么可以替代 queue.BytesQueue ?
	// chunks [][]byte替换queue.BytesQueue
	// RingBuffer

	lock        sync.RWMutex
	entryBuffer []byte
	onRemove    onRemoveCallback

	isVerbose    bool
	statsEnabled bool
	//logger       Logger
	//clock        clock
	lifeWindow   uint64

	hashmapStats map[uint64]uint32
	stats        Stats
}

// BytesQueue is a non-thread safe queue type of fifo based on bytes array.
// For every push operation index of entry is returned.
// It can be used to read the entry later
type BytesQueue struct {
	full         bool
	array        []byte
	capacity     int
	maxCapacity  int
	head         int
	tail         int
	count        int
	rightMargin  int
	headerBuffer []byte
	verbose      bool
}


// Stats stores cache statistics
type Stats struct {
	// Hits is a number of successfully found keys
	Hits int64 `json:"hits"`
	// Misses is a number of not found keys
	Misses int64 `json:"misses"`
	// DelHits is a number of successfully deleted keys
	DelHits int64 `json:"delete_hits"`
	// DelMisses is a number of not deleted keys
	DelMisses int64 `json:"delete_misses"`
	// Collisions is a number of happened key-collisions
	Collisions int64 `json:"collisions"`
}


type onRemoveCallback func(wrappedEntry []byte, reason RemoveReason)

// RemoveReason is a value used to signal to the user why a particular key was removed in the OnRemove callback.
type RemoveReason uint32


// Config for BigCache
type Config struct {
	// Number of cache shards, value must be a power of two
	Shards int
	// Time after which entry can be evicted
	LifeWindow time.Duration
	// Interval between removing expired entries (clean up).
	// If set to <= 0 then no action is performed. Setting to < 1 second is counterproductive — bigcache has a one second resolution.
	CleanWindow time.Duration
	// Max number of entries in life window. Used only to calculate initial size for cache shards.
	// When proper value is set then additional memory allocation does not occur.
	MaxEntriesInWindow int
	// Max size of entry in bytes. Used only to calculate initial size for cache shards.
	MaxEntrySize int
	// StatsEnabled if true calculate the number of times a cached resource was requested.
	StatsEnabled bool
	// Verbose mode prints information about new memory allocation
	Verbose bool
	// Hasher used to map between string keys and unsigned 64bit integers, by default fnv64 hashing is used.
	Hasher Hasher
	// HardMaxCacheSize is a limit for cache size in MB. Cache will not allocate more memory than this limit.
	// It can protect application from consuming all available memory on machine, therefore from running OOM Killer.
	// Default value is 0 which means unlimited size. When the limit is higher than 0 and reached then
	// the oldest entries are overridden for the new ones.
	HardMaxCacheSize int
	// OnRemove is a callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	// ignored if OnRemoveWithMetadata is specified.
	OnRemove func(key string, entry []byte)
	// OnRemoveWithMetadata is a callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A structure representing details about that specific entry.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	OnRemoveWithMetadata func(key string, entry []byte, keyMetadata Metadata)
	// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A constant representing the reason will be passed through.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	// Ignored if OnRemove is specified.
	OnRemoveWithReason func(key string, entry []byte, reason RemoveReason)

	onRemoveFilter int

	// Logger is a logging interface and used in combination with `Verbose`
	// Defaults to `DefaultLogger()`
	Logger Logger
}