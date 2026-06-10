```GO
type Span interface{
	// Span执行单元结束
	Finish()
	// 带结束时间和日志记录列表信息，Span执行单元结束；也就是说结束时间可以业务指定; 日志列表也可以直接添加； 目前还不知道这个的使用场景；
	FinishWithOptions(opts FinishOptions)
	// 把Span的Baggage封装成SpanContext
	Context() SpanContext
	// 设置Span的操作名称
	SetOperationName(operationName string) Span
	// 设置Span Tag
	SetTag(key string, value interface{}) Span
	// 设置Span的log.Field列表；
	// span.LogFields(
	// 		log.String("event", "soft error"),
	// )
	LogFields(fields ...log.Field) 
	// key-value列表={"event": "soft error", "type": "cache timeout", "waited.millis":1500}
	LogKV(alternatingKeyValues ...interface{})
	// LogFields与LogKV类似，只是前者已封装好；
	
	// 设置span的Baggage：key-value， 用于跨进程上下文传输
	SetBaggageItem(restrictedKey, value string) Span
	// 通过key获取value；
	BaggageItem(restrictedKey string) string
	// 获取Span所在的调用链tracer
	Tracer() Tracer
	// 废弃，改用LogFields 或者 LogKV
	LogEvent(event string)
	// 同上
	LogEventWithPayload(event string, payload interface{})；
	// 同上
	Log(data LogData)
}
```

## SpanContext

首先，需要看一小段代码:

```shell
package main

import "fmt"

type Fun struct{}

func main() {
    var fun1, fun2 = Fun{}, Fun{}
    if fun1 == fun2 {
        fmt.Println("fun1==fun2")
    } else {
        fmt.Println("fun1!=fun2")
    }
}

// 执行结果：fun1==fun2
// 比较两个值是否相等，取决于：值和类型，都相等这表示相同；
```

Context用于上下文数据传输使用，在OpenTracing标准中，Span之间跨进程调用时，会使用SpanContext传输Baggage携带信息。通过context标准库实现；如：

```shell
type contextKey struct{}

var activeSpanKey = contextKey{}

// 封装span到context中
func ContextWithSpan(ctx context.Context, span Span) context{
	return context.WithValue(ctx, activeSpanKey, span)
}

// 从ctx中通过activeSpanKey取出Span，这里可以看到不同服务的activeSpanKey值，是相同的，上面的DEMO可以说明。
func SpanFromContext(ctx context.Context) Span{
	val:=ctx.Value(activeSpanKey)
	if sp, ok := val.(Span); ok{
		return sp
	}
	return nil
}
```

通过context上下文的activeSpanKey，我们可以获得Span，并创建新的span，如：

```shell
func startSpanFormContextWithTracer(ctx context.Context, tracer Tracer, operationName string, opts ...StartSpanOption) (Span, context.Context){
	// 首先从上下文看是否能够获取到span，如果获取不到，再创建tracer和span；
	if parentSpan:= SpanFromContext(ctx); parentSpan !=nil {
		opts = append(opts, ChildOf(parentSpan.Context()))
	}
	
	span := tracer.StartSpan(operationName, opts...)
	return span, ContextWithSpan(ctx, span)
}
```

### Jaeger 是如何实现跨进程的？

# Propagation
1. 跨进程的trace信息传输通过SpanContext传递baggage信息，把 baggage 存储到 context 上下文中，
则需要key:value, 这个key决定了value值和value类型；
2. key:value存储到context中，需要借助于读写操作；`notice: 这个借鉴了io.Reader和io.Writer等思想，通过组合模式使得实现变得更加灵活`
3. 目前支持三种key值，对应三种value值类型，
分别是：byte流、TextMap和http.Header；其中后两者都是map结构；只是RPC或者Web时，用http.Header，其他使用TextMap或者byte流；

4. TextMapWriter 
5. `type TextMapCarrier map[string]string`



