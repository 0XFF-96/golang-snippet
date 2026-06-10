package client

// Sampler decides whether a new trace should be sampled or not.
type Sampler interface {
	// IsSampled decides whether a trace with given `id` and `operation`
	// should be sampled. This function will also return the tags that
	// can be used to identify the type of sampling that was applied to
	// the root span. Most simple samplers would return two tags,
	// sampler.type and sampler.param, similar to those used in the Configuration
	IsSampled(id TraceID, operation string) (sampled bool, tags []Tag)

	// Close does a clean shutdown of the sampler, stopping any background
	// go-routines it may have started.
	Close()

	// Equal checks if the `other` sampler is functionally equivalent
	// to this sampler.
	// TODO remove this function. This function is used to determine if 2 samplers are equivalent
	// which does not bode well with the adaptive sampler which has to create all the composite samplers
	// for the comparison to occur. This is expensive to do if only one sampler has changed.
	Equal(other Sampler) bool
}


// Tag is a simple key value wrapper.
// TODO deprecate in the next major release, use opentracing.Tag instead.
type Tag struct {
	key   string
	value interface{}
}