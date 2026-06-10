### 什么是 groupcache ?

```group cache 的特点是
groupcache coordinates
      cache fills such that only one load in one process of an entire
      replicated set of processes populates the cache, then multiplexes
      the loaded value to all callers.
```


### group cache 的特点

### 相关
1、分布式缓存库
2、数据无版本概念, 也就是一个key对应的value是不变的，并没有update
3、节点之间可互访，自动复制热点数据

### 实现

按Groupcache项目页面上的宣称的替代 Memcached 的大多数应用场景表示不大现实，
GroupCache在我看来基本上只能适用于静态数据，比如下载服务中的文件缓存等。 



### singleflight 的相关实现
1. 这里的实现和 go-zero 中 SharedCalls 的实现本质是一样的。 
2. 而且 SharedCall 将名字换了一下和重构了一下代码，就拿出来了...
用 defer() 机制

目的是
```
	// SharedCalls lets the concurrent calls with the same key to share the call result.
	// For example, A called F, before it's done, B called F. Then B would not execute F,
	// and shared the result returned by F which called by A.
	// The calls with the same key are dependent, concurrent calls share the returned values.
	// A ------->calls F with key<------------------->returns val
	// B --------------------->calls F with key------>returns val
``` 