[啤酒]技术总监毛剑精选 Go 预习资料

Quick Start：
https://golang.org/doc/install
https://tour.golang.org/welcome/1

语言层面：
《Effective Go》 https://golang.org/doc/effective_go.html （我是看这个入门，官方推荐的文档，轻松易读
https://item.jd.com/12187988.html （比较详细的书籍

实战 & 规范：

《Code Review Comments》
en：https://github.com/golang/go/wiki/CodeReviewComments
cn：https://studygolang.com/articles/10960

《High Performance Go Workshop》
https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#overview

《Practical Go: Real world advice for writing maintainable Go programs》
https://dave.cheney.net/practical-go/presentations/qcon-china.html

Runtime：
https://draveness.me/golang/ （类似书籍比较多，还有一本雨痕的可以找找

官方文档：
golang.org
godoc.org
=================

其他国内的书我看的比较少，最推荐上面的Effective Go，另外了解下B站的框架：
https://github.com/bilibili/kratos


[啤酒]Go 进阶训练营免费福利资料

链接:https://pan.baidu.com/s/1O3jKZ-qvXnWkJbaupoVlmQ 密码:33lt
Go 并发编程 140 页干货 PPT：
链接: https://pan.baidu.com/s/1iWH3_IXaTwwBcTl4h3kz-w 密码: 61ys
大厂软件开发案例：https://con.infoq.cn/archives/?conf=qcons
1、https://www.yuque.com/docs/share/6bb19a9c-99d3-414b-bd69-bdb788415b30?# 

### 相关？

1、go 资深工程师课程-配套文档
2、企业级网关系统相关
3、秒杀-相关
4、rocksdb + raft 做 kv 系统相关
5、需要知道后端一些重要常量
6、缓存、存储、队列 三个最重要的中间价之一。
7、7 days- https://github.com/geektutu/7days-golang  
8、作为一个中间件：https://github.com/nikoksr/notify  
9、定时查看 trending  看看别人在做一些什么工作


### 播放历史系统设计


1、整体架构（为什么需要这样的设计） 

2、网卡的限制， redis qps 单节点 8 w。 
风险（我们同样存这个风险）： 
这里存在两个风险: 
1、history-service 重启过程中，预聚合的消息丢失；2、history-job 读取 redis 构建数据，但 redis 丢失；
我们在这里进行了 trade-off，高收敛比的设计，意味着存在数据丢失的风险，对于历史场景，非 L0 的业务服务/数据，我们认为极端情况下可接受。 

3、 API Gateway 的流量都会触发高频的 per-rpc auth。 






