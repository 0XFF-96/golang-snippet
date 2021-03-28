### 微服务注册中心文章系列
1. [比较系统介绍 consul 微服务的实现] https://laravelacademy.org/post/21213
2. 利用 etcd 实现一个简单的注册中心源代码： https://github.com/awei0217/go_common/blob/master/grpc/helloworld_demo/register_center/etcd_resolver.go 
3. https://segmentfault.com/a/1190000020944777 

### gRPC服务发现&负载均衡
- https://segmentfault.com/a/1190000008672912


### grpc 源码解读
1. https://learnku.com/articles/35336

```
http://www.grpc.io/docs/
https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
```

负载均衡和服务发现的测试方法
- 运行3个服务端S1、S2、S3，1个客户端C，观察各服务端接收的请求数是否相等？
- 关闭1个服务端S1，观察请求是否会转移到另外2个服务端？
- 关闭1个服务端S1，观察请求是否会转移到另外2个服务端？
- 关闭Etcd3服务器，观察客户端与服务端通信是否正常？
- 重新启动Etcd3服务器，服务端上下线可自动恢复正常。
- 关闭所有服务端，客户端请求将被阻塞。


### go-zero 如何实现负载均衡？
- https://zhuanlan.zhihu.com/p/332863487
- 

### 需要区分的概念
- 在哪个层面上进行 [负载均衡] ？ 各有什么好处？
- 如何区分不同组件作为 provider、consumer、discover ？
- 常见后端组件互联互通网络架构？（一个组件如何调用另一个组件的？）


### 参考的库
- https://github.com/wwcd/grpc-lb
- https://github.com/wothing/wonaming

### FAQ 
- 想 eruaka 和 b 站的 discovery 方案和 B 站的方案是属于同一类型的吗？ 它们之间的优缺点是什么？（AP？）（CAP ?)
- 