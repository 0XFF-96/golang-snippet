package client


// SpanAllocator abstraction of managign span allocations
type SpanAllocator interface {
	Get() *Span
	Put(*Span)
}

type syncPollSpanAllocator struct {
	spanPool sync.Pool
}
