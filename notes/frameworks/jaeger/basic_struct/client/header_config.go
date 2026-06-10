package client


// HeadersConfig contains the values for the header keys that Jaeger will use.
// These values may be either custom or default depending on whether custom
// values were provided via a configuration.
type HeadersConfig struct {
	// JaegerDebugHeader is the name of HTTP header or a TextMap carrier key which,
	// if found in the carrier, forces the trace to be sampled as "debug" trace.
	// The value of the header is recorded as the tag on the root span, so that the
	// trace can be found in the UI using this value as a correlation ID.
	JaegerDebugHeader string `yaml:"jaegerDebugHeader"`

	// JaegerBaggageHeader is the name of the HTTP header that is used to submit baggage.
	// It differs from TraceBaggageHeaderPrefix in that it can be used only in cases where
	// a root span does not exist.
	JaegerBaggageHeader string `yaml:"jaegerBaggageHeader"`

	// TraceContextHeaderName is the http header name used to propagate tracing context.
	// This must be in lower-case to avoid mismatches when decoding incoming headers.
	TraceContextHeaderName string `yaml:"TraceContextHeaderName"`

	// TraceBaggageHeaderPrefix is the prefix for http headers used to propagate baggage.
	// This must be in lower-case to avoid mismatches when decoding incoming headers.
	TraceBaggageHeaderPrefix string `yaml:"traceBaggageHeaderPrefix"`
}
