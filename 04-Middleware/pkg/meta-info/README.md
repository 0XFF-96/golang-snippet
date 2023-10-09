### 元信息传递

1、通过 http header 
2、通过 grpc metadata outgoing 和 incoming 两类


### Concept 
1、该库提供了一种在 go 语言的 `context.Context` 中保存用于跨服务传递的元信息的统一接口。

from bytedance/pkg metainfo 
```
根据数据传递的场景，元信息被分为两种类别，*transient* 和 *persistent* —— 前者只会传递一跳，从客户端传递到下游服务端，然后消失；后者需要在整个服务调用链上一直传递，直到被丢弃。

由于传递过程使用了 go 语言的 `context.Context`，因此，为了避免服务端的 `Context` 实例直接传递给用于调用下游服务的客户端时造成从更上游服务传递来的 transient 数据被继续透传，需要引入一个中间态：*transient-upstream*，以和当前服务自己设置的 transient 数据作区分。该类别仅用于实现 transient 的语义，在接口上，transient 和 transient-upstream 并没有区别。

该库被设计成针对 `context.Context` 进行操作的接口集合，而具体的元数据在网络上传输的形式和方式，应由支持该库的框架来实现。通常，终端用户不应该关注其具体传输形式，而应该仅依赖该库提供的抽象接口。

```

### REFERENCE 
