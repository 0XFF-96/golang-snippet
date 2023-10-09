package client

import (
	"github.com/opentracing/opentracing-go"
)

type Observer interface {
	OnStartSpan(operationName string, options opentracing.StartSpanOptions) SpanObserver
}

// SpanObserver is created by the Observer and receives notifications about
// other Span events.
//
// Deprecated: use jaeger.ContribSpanObserver instead.
type SpanObserver interface {
	OnSetOperationName(operationName string)
	OnSetTag(key string, value interface{})
	OnFinish(options opentracing.FinishOptions)
}



// ContribObserver can be registered with the Tracer to receive notifications
// about new Spans. Modelled after github.com/opentracing-contrib/go-observer.
type ContribObserver interface {
	// Create and return a span observer. Called when a span starts.
	// If the Observer is not interested in the given span, it must return (nil, false).
	// E.g :
	//     func StartSpan(opName string, opts ...opentracing.StartSpanOption) {
	//         var sp opentracing.Span
	//         sso := opentracing.StartSpanOptions{}
	//         if spanObserver, ok := Observer.OnStartSpan(span, opName, sso); ok {
	//             // we have a valid SpanObserver
	//         }
	//         ...
	//     }
	OnStartSpan(sp opentracing.Span, operationName string, options opentracing.StartSpanOptions) (ContribSpanObserver, bool)
}
