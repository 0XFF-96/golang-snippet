### Grpc 源码相关
// gRPC 有两个方法
// 将信息放到头部的 md 
// 这条链路的全过程是怎样子的？
// 探究一下
在 gRPC 文档中：
Clients may use
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)
to store tokens and other authentication-related data. To gain access to the
`metadata.MD` object, a server may use
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext).
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New).