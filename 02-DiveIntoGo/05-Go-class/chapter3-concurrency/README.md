### 外部资料
1. http://github.com/da440dil/go-workgroup 
2.

### NOTE
1. 一个 gorountine 应该是完全由调用者管理的
2. 

### Golang 内存模型【✅】
1. https://golang.org/ref/mem
2. 什么是内存模型？
3. Memory Reordering
4. 什么是 memory barrier / fence ? 所有对内存对写操作，都要进行扩散
5. https://www.jianshu.com/p/5e44168f47a3


### Mutex
1. mutex 的三个实现方式？

### 如何使用 errgroup 
1. https://github.com/go-kratos/kratos/tree/master/pkg/sync/errgroup

### sync.Pool 
1. 放进 Pool 中的对象，会在说不准什么时候被回收掉。
所以如果事先 Put 进去 100 个对象，下次 Get 的时候发现 Pool 是空也是有可能的 ?
因为 Pool 自己有 GC 回收机制

2. 


### Chan 同步模型
1、Timing out
2、Moving on
3、Pipeline、streaming
4、Fan-out, Fan-in
5、Cancellation
6、Close 先于 Receive 发生(类似 Buffered)。
    不需要传递数据，或者传递 nil。
    非常适合去掉和超时控制。
7、Contex

https://blog.golang.org/concurrency-timeouts
https://blog.golang.org/pipelines
https://talks.golang.org/2013/advconc.slide#1
https://github.com/go-kratos/kratos/tree/master/pkg/sync
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html

```
If any given Send on a channel CAN cause the sending goroutine to block:
    Not allowed to use a Buffered channel larger than 1.
    Buffers larger than 1 must have reason/measurements.
    Must know what happens when the sending goroutine blocks.
If any given Send on a channel WON’T cause the sending goroutine to block:
    You have the exact number of buffers for each send.
        Fan Out pattern
    You have the buffer measured for max capacity.
        Drop pattern
Less is more with buffers.
    Don’t think about performance when thinking about buffers.
    Buffers can help to reduce blocking latency between signaling.
    Reducing blocking latency towards zero does not necessarily mean better throughput.
    If a buffer of one is giving you good enough throughput then keep it.
Question buffers that are larger than one and measure for size.
    Find the smallest buffer possible that provides good enough throughput.
```

### 利用 Context 染色
1、染色，API 重要性，Trace
2、https://github.com/go-kratos/kratos/blob/master/pkg/net/metadata/key.go

### REFERENCE 

https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html
https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html
https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html
https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency
https://golang.org/ref/mem
https://blog.csdn.net/caoshangpa/article/details/78853919
https://blog.csdn.net/qcrao/article/details/92759907
https://cch123.github.io/ooo/
https://blog.golang.org/codelab-share
https://dave.cheney.net/2018/01/06/if-aligned-memory-writes-are-atomic-why-do-we-need-the-sync-atomic-package
http://blog.golang.org/race-detector
https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races
https://www.ardanlabs.com/blog/2014/06/ice-cream-makers-and-data-races-part-ii.html
https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549
https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c
https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268
https://medium.com/a-journey-with-go/go-ordering-in-select-statements-fd0ff80fd8d6
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
https://www.ardanlabs.com/blog/2013/10/my-channel-select-bug.html
https://blog.golang.org/io2013-talk-concurrency
https://blog.golang.org/waza-talk
https://blog.golang.org/io2012-videos
https://blog.golang.org/concurrency-timeouts
https://blog.golang.org/pipelines
https://www.ardanlabs.com/blog/2014/02/running-queries-concurrently-against.html
https://blogtitle.github.io/go-advanced-concurrency-patterns-part-3-channels/
https://www.ardanlabs.com/blog/2013/05/thread-pooling-in-go-programming.html
https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html
https://blogtitle.github.io/categories/concurrency/
https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
https://blog.golang.org/context
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
https://golang.org/ref/spec#Channel_types
https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view
https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
https://blog.golang.org/context
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
https://golang.org/doc/effective_go.html#concurrency
https://zhuanlan.zhihu.com/p/34417106?hmsr=toutiao.io
https://talks.golang.org/2014/gotham-context.slide#1s
https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39

