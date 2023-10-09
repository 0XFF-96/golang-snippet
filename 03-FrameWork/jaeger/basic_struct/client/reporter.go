package client


// Reporter is called by the tracer when a span is completed to report the span to the tracing collector.
type Reporter interface {
	// Report submits a new span to collectors, possibly asynchronously and/or with buffering.
	Report(span *Span)

	// Close does a clean shutdown of the reporter, flushing any traces that may be buffered in memory.
	Close()
}
