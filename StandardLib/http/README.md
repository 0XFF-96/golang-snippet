### 相关术语
1.  shallow clone 浅拷贝
2. 



### 核心方法
1. RoundTrip 
```
// RoundTripper is an interface representing the ability to execute a
// single HTTP transaction, obtaining the Response for a given Request.
//
// A RoundTripper must be safe for concurrent use by multiple
// goroutines.
type RoundTripper interface {
	// RoundTrip executes a single HTTP transaction, returning
	// a Response for the provided Request.
	//
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// A non-nil err should be reserved for failure to obtain a
	// response. Similarly, RoundTrip should not attempt to
	// handle higher-level protocol details such as redirects,
	// authentication, or cookies.
	//
	// RoundTrip should not modify the request, except for
	// consuming and closing the Request's Body. RoundTrip may
	// read fields of the request in a separate goroutine. Callers
	// should not mutate or reuse the request until the Response's
	// Body has been closed.
	//
	// RoundTrip must always close the body, including on errors,
	// but depending on the implementation may do so in a separate
	// goroutine even after RoundTrip returns. This means that
	// callers wanting to reuse the body for subsequent requests
	// must arrange to wait for the Close call before doing so.
	//
	// The Request's URL and Header fields must be initialized.
	RoundTrip(*Request) (*Response, error)
}
```

2. 如何 skip test in short mode ?
```
	if testing.Short() {
		t.Skip("skipping test in -short mode")
	}
```

有点意思


3. 重点二： readLoop 和 writeLoop 

```
	// Verify no outstanding requests after readLoop/writeLoop
	// goroutines shut down.
```


### 为什么没有 【连接池】不能被服用？
- 如何监控这种情况？
- 如何知道当前存储的连接是多少？

- http://xiaorui.cc/archives/7172



### 为什么 resp.Body.Close() 为什么需要 close ?

- 如果不 close 的话，会出现连接泄漏（表面原因） 
- 根本原因：？

ref: https://segmentfault.com/a/1190000020086816?utm_source=sf-similar-article

综上分析，可以发现，readLoop和 writeLoop 两个goroutine在 写入请求并获取response返回后，
并没有跳出for循环，而继续阻塞在 下一次for循环的select 语句里面，
所以，两个函数所在的goroutine并没有运行结束，
导致了最开的现象: goroutine持续增加导致内存持续增加。 


Close 的时候发生了什么？

close的主要逻辑是通过调用 readLoop 的第89行定义的earlyCloseFn 方法， 向 waitForBodyRead 的chan写入false，进而让 readLoop 退出阻塞，终止 readLoop 的 goroutine

readLoop 退出的时候，关闭 closech chan，进而让 writeLoop 退出阻塞，终止 writeLoop 的goroutine


