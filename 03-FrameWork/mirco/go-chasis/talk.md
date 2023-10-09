### talk notes

go-chassis 在 shoppe  
1、之前是 python gunicorn django , 后面转 go  

2、技术栈 redis, mysql, elastic, mesos, ceph, redis， kafka 

3、shoppe: grpc/ http 服务， 业务 api gateway, 定时/异步任务， 异构数据平台， id 生成服务， 多级缓存。  

4、统一业务框架， 服务治理规范， 统一监控规范， 全链路追踪， 服务注册/服务发现 

5、没有客户端负载均衡？
6、目前的服务状态， 服务提供方、服务使用方、Service-center, etcd. 注册及定时心跳， 长链接实时推送， 定时拉去服务元数据，

 7、cat 日志客户端？apollo ( 携程那一套） 

8、配置热更新， go-archaius, 热更新支持， 支持多配置源， 支持自定义配置源。  

9、go-archaius + apollo 作为配置源接入。  

10、cat 美团开源的实时应用监控平台。 opentracing 。 https://github.com/dianping/cat 

11、异步批量上报， 上报上下文导致频繁的内存分配：sync.Pool。   

12、100 w/秒的日志 api, 异步， fmt.Sprintf 很慢， 取文件行号， 利用 RetAddr 做缓存， 时间转字符串也有开销， 缓存到秒。 

 13、sync.Pool + slice 与内存泄漏的相关问题。 

 14、数据库的问题， tx 意外的写入成功，

   15、多级缓存。   不太懂，为什么写入不成功。 file:///Users/lijinrui/Downloads/2.2.5%20go-Chassis%20%E5%9C%A8%20shopee%20%E4%BE%9B%E5%BA%94%E9%93%BE%E7%9A%84%E5%AE%9E%E8%B7%B5%20(1).pdf 。 

  https://github.com/golang/go/issues/15123 1、什么是 prepare stat ?