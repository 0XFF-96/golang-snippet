### 什么是 grpc 命名空间冲突？

1、https://developers.google.com/protocol-buffers/docs/reference/go/faq#namespace-conflict

### ResourceExhausted desc

1、rpc error: code = ResourceExhausted desc 
= grpc: received message larger than max

2、什么原因导致这个 error ， 有什么方法解决

### grpc 的 load-balance 在 k8s 环境下失效问题

1、https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/ [done]
2、https://www.practicalnetworking.net/series/packet-traveling/packet-traveling/ 【网络】


### grpc 处理请求源码

1、grpc 每调用一次，都会用一个协程处理
2、在协程里面开启协程，如果没有保护的话，你后面的操作就可能遇到 context cancel 的报错，从而导致事件发送不成功。 
3、event_center 的相关依赖问题。是怎么进行调用的？
4、追踪一下 grpc 的源码的调用过程。 

