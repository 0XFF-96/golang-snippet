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


### grpc 是如何进行超时传递的？
1、超时传递原理：https://xiaomi-info.github.io/2019/12/31/grpc-custom-ns/ 
2、https://xiaomi-info.github.io/2019/12/31/grpc-custom-ns/

### 什么是 in-comming context 和 什么是 out-going context 
1. 它们之间的基本原来是什么？
2.  in-comming 和 out-going context 之间的转换过程
3. 这两个之间应该什么时候用？https://stackoverflow.com/questions/57060602/what-is-the-difference-between-metadata-fromoutgoingcontext-and-metadata-frominc 
4. 官方文档的相关用法：https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md
