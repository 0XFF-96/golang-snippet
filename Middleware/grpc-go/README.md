### grpc-go 源码解构

### FAQ 
1、IncomingContext 和 OutgoingContext 之间有什么区别，他们是如何被传递下去的？
2、grpc 的 deadline context 是如何跨服务传递的？

### gRPC 的相关例子
1. https://github.com/vladimirvivien/go-grpc
2. https://www.bookstack.cn/read/grpc-read/12-grpc%20%E6%95%B0%E6%8D%AE%E6%B5%81%E8%BD%AC.md
grpc 源码剖析: https://github.com/lubanproj/grpc-read
3.【小米】context canceled 报错 https://xiaomi-info.github.io/2019/12/30/grpc-deadline/

### 重点函数
1. 构造调用请求 `newClientStream 源码`, 是微服务相互调用的基础
2. ```createHeaderFields``` 是 grpc 构造 header Frames 的源码
  ```(d *decodeState) processHeaderField(f hpack.HeaderField)``` 是 grpc 服务解析的地方
