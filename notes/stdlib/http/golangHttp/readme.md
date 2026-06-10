### conn 在不同时期的转换过程

wantedConn -> persistConn -> tcpConn 

### pool or queue 有几种？

### 疑惑

1. 为什么需要 roundTripRequest 到 connectMethod 的转换？
```
connectMethodForRequest(treq *transportRequest) (cm connectMethod)
```


2. queueForIdleConn(w *wantConn) 函数中
// Look for most recently-used idle connection.  为什么需要这个策略？


3. 关于 http1 和 http2 的区别，在 net/http 包的体现

alt roundTrip , 有很多分支是基于 alt == nil 判断。
- pconn 在被使用中时，是否应该移除 idle conn 池
```
					// HTTP/2: multiple clients can share pconn.
					// Leave it in the list.
```

```
					// HTTP/1: only one client can use pconn.
					// Remove it from the list.
```

4. 当 maxConnPerHost 满了的时候，谁 dial For 呢？
这个方法
```
// decConnsPerHost decrements the per-host connection count for key,
// which may in turn give a different waiting goroutine permission to dial.
func (t *Transport) decConnsPerHost(key connectMethodKey) {
```
