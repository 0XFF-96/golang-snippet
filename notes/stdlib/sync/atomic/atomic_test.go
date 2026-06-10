package atomic

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

// The following different-struct-memory shows how to use Value for periodic program config updates
// and propagation of the changes to worker goroutines.
func TestExampleValue_config(t *testing.T) {
	var config atomic.Value // holds current server configuration
	// Create initial config value and store into config.
	config.Store(loadConfig())
	go func() {
		// Reload config every 10 seconds
		// and update config value with the new version.
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				// Handle request r using config c.
				_, _ = r, c
			}
		}()
	}
}

// 为什么减一, 需要这样实现？
func TestDeltaOneUint32(t *testing.T) {
	var i uint32
	i = 30
	atomic.AddUint32(&i, ^uint32(0))
	require.Equal(t, i, uint32(29))

	i = 0
	atomic.AddUint32(&i, ^uint32(0))
	// 溢出了
	// 4294967295
	t.Log(i)
}

func TestDeltaInt32(t *testing.T) {
	var i int32
	i = 30
	atomic.AddInt32(&i, ^int32(0))
	require.Equal(t, i, int32(29))

	i = 0
	atomic.AddInt32(&i, ^int32(0))
	require.Equal(t, i, int32(-1))
}
