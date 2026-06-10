package remote

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jaegertracing/jaeger-client-go"
)

const (
	defaultHostPort        = "localhost:5778"
	defaultRefreshInterval = time.Second * 5
)

// Option is a function that sets some option on the Throttler
type Option func(options *options)

// Options is a factory for all available options
var Options options

type options struct {
	metrics                   *jaeger.Metrics
	logger                    jaeger.Logger
	hostPort                  string
	refreshInterval           time.Duration
	synchronousInitialization bool
}

// Metrics creates an Option that initializes Metrics on the Throttler, which is used to emit statistics.
func (options) Metrics(m *jaeger.Metrics) Option {
	return func(o *options) {
		o.metrics = m
	}
}

// Logger creates an Option that sets the logger used by the Throttler.
func (options) Logger(logger jaeger.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// HostPort creates an Option that sets the hostPort of the local agent that keeps track of credits.
func (options) HostPort(hostPort string) Option {
	return func(o *options) {
		o.hostPort = hostPort
	}
}

// RefreshInterval creates an Option that sets how often the Throttler will poll local agent for
// credits.
func (options) RefreshInterval(refreshInterval time.Duration) Option {
	return func(o *options) {
		o.refreshInterval = refreshInterval
	}
}

// SynchronousInitialization creates an Option that determines whether the throttler should synchronously
// fetch credits from the agent when an operation is seen for the first time. This should be set to true
// if the client will be used by a short lived service that needs to ensure that credits are fetched upfront
// such that sampling or throttling occurs.
func (options) SynchronousInitialization(b bool) Option {
	return func(o *options) {
		o.synchronousInitialization = b
	}
}

func applyOptions(o ...Option) options {
	opts := options{}
	for _, option := range o {
		option(&opts)
	}
	if opts.metrics == nil {
		opts.metrics = jaeger.NewNullMetrics()
	}
	if opts.logger == nil {
		opts.logger = jaeger.NullLogger
	}
	if opts.hostPort == "" {
		opts.hostPort = defaultHostPort
	}
	if opts.refreshInterval == 0 {
		opts.refreshInterval = defaultRefreshInterval
	}
	return opts
}

// GetJSON makes an HTTP call to the specified URL and parses the returned JSON into `out`.
func GetJSON(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	return ReadJSON(resp, out)
}

// ReadJSON reads JSON from http.Response and parses it into `out`
func ReadJSON(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	if out == nil {
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

