package remote

import (
	"testing"
	"time"

	"github.com/jaegertracing/jaeger-client-go"
	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	options := applyOptions()
	assert.Equal(t, "localhost:5778", options.hostPort)
	assert.Equal(t, time.Second*5, options.refreshInterval)
	assert.NotNil(t, options.metrics)
	assert.NotNil(t, options.logger)
	assert.False(t, options.synchronousInitialization)
}

func TestOptions(t *testing.T) {
	metrics := jaeger.NewNullMetrics()
	logger := jaeger.NullLogger
	options := applyOptions(
		Options.Metrics(metrics),
		Options.Logger(logger),
		Options.HostPort(":"),
		Options.RefreshInterval(time.Second),
		Options.SynchronousInitialization(true),
	)
	assert.Equal(t, ":", options.hostPort)
	assert.Equal(t, time.Second, options.refreshInterval)
	assert.Equal(t, metrics, options.metrics)
	assert.Equal(t, logger, options.logger)
	assert.True(t, options.synchronousInitialization)
}
