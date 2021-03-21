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