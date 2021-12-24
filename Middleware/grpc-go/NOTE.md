# gRPC Notes 

### Outline 大纲


### rpcx 框架

1、https://doc.rpcx.io/part3/timeout.html 。 值得研究一下学习。 
2、https://doc.rpcx.io/
3、https://github.com/smallnest/rpcx

### go-chasis 的 Invoker 抽象和 go-micro 的关系
1、

### gRPC 底层原理

1、http://cnblogs.com/sunsky303/p/11119300.html

### gRPC 客户端

1、https://forum.rentsoft.cn/thread/2 


### gRPC Gateway & Proxy 

1、https://github.com/fullstorydev/grpcurl
2、https://github.com/mercari/grpc-http-proxy
3、https://github.com/gdong42/grpc-mate
4、https://github.com/grpc-ecosystem/grpc-gateway

### gRPC 和 K8s 平滑发布问题

1、因为基于 k8s 的 service 的 DNS 刷新时间有间隔，所有不能做到平滑更新和平滑发布。 （发布的一段时间，会有很多错误，采用 rolling update 的时候） 
2、grpc 怎么配置多个双向认证的证书来进行细粒度的权限认证？
3、为什么移动端会连不上 http2.0, grpc 不是不能直接通过 http 请求访问的吗？ grcp-gateway 的方式是指什么？ grpc 怎么分别提供 http1.1 和 http2.0 相关的服务。 
【grpc 公网 ppt 】


### grpc Log 问题

Logrus can also be made as a backend for gRPC library internals. 
For that use ReplaceGrpcLogger.
Server Interceptor
Below is a JSON formatted example of a log that would be logged by the server interceptor:

这个是从哪里来的 ？ log_peer.address:172.20.0.129:37320， 

是啥意思？peer.address 对等地址。 

### gRPC 如何连通 client ?

1、为什么使用 serviceName:port , 能够连通其他的服务。 这是什么原理？
2、为什么 127.0.0.1:port, 这一种联通方式同样 ok ？
3、GetEndPoint()   
4、name resolution , https://github.com/grpc/grpc/blob/master/doc/naming.md 。 
5、http://doc.rpcx.io  
6、有没有什么办法可以在本地，调用 staging 的 内部服务？



### Pending
1. 学一下使用 grpc 的 channelz.IsOn()
2. grpc Service-config 设计文档， https://github.com/grpc/grpc/blob/master/doc/service_config.md
3. 

