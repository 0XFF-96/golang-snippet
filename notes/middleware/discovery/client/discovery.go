package client

import (
	"context"
	"math/rand"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	http "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// Discovery is discovery client.
type Discovery struct {
	c          *Config
	once       sync.Once
	ctx        context.Context
	cancelFunc context.CancelFunc
	httpClient *http.Client

	node    atomic.Value
	nodeIdx uint64

	mutex       sync.RWMutex
	apps        map[string]*appInfo
	registry    map[string]struct{}
	lastHost    string
	cancelPolls context.CancelFunc

	delete chan *appInfo
}

// Config discovery configures.
type Config struct {
	Nodes  []string
	Region string
	Zone   string
	Env    string
	Host   string
}

// Reload reload the config
func (d *Discovery) Reload(c *Config) {
	fixConfig(c)
	d.mutex.Lock()
	d.c = c
	d.mutex.Unlock()
}

// Close stop all running process including discovery and register
func (d *Discovery) Close() error {
	return nil
}

// Register Register an instance with discovery and renew automatically
func (d *Discovery) Register(ins *Instance) (cancelFunc context.CancelFunc, err error) {
	return nil, nil
}

// register Register an instance with discovery
func (d *Discovery) register(ctx context.Context, ins *Instance) (err error) {
	return nil
}

// renew Renew an instance with discovery
func (d *Discovery) renew(ctx context.Context, ins *Instance) (err error) {
	return
}

// cancel Remove the registered instance from discovery
func (d *Discovery) cancel(ins *Instance) (err error) {
	return
}

// Set set ins status and metadata.
func (d *Discovery) Set(ins *Instance) error {
	return nil
}

// set set instance info with discovery
func (d *Discovery) set(ctx context.Context, ins *Instance) (err error) {
	return
}

func (d *Discovery) serverproc() {

}

func (d *Discovery) tryAppendSeedNodes() {
}

func (d *Discovery) pickNode() string { return "" }

func (d *Discovery) switchNode() {
}

func (d *Discovery) polls(
	ctx context.Context,
	) (apps map[string]*InstancesInfo, err error) {
	return nil, nil
}

// 这个 broadcast 的操作
// 和
func (d *Discovery) broadcast(apps map[string]*InstancesInfo) {
}

func (d *Discovery) newParams(c *Config) url.Values {
	return nil
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}

	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// Shuffle really ought not be called with n that doesn't fit in 32 bits.
	// Not only will it take a very long time, but with 2³¹! possible permutations,
	// there's no way that any PRNG can have a big enough internal state to
	// generate even a minuscule percentage of the possible permutations.
	// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

