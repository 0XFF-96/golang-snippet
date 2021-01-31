### F&A
1）如何理解服务发现这一概念？
2）服务发现要解决的问题是什么？
3）Eureka是如何设计服务发现？
4）Eureka如何保证高可用与数据一致？

### 核心流程 (client 端)
1、客户端发起服务注册；
2、服务端保存信息至注册表；
3、客户端定时发送心跳请求；
4、服务端服务剔除与自我保护
5、客服端发送服务下线请求
6、客户端获取服务端注册表；
7、客户端整合服务发现

### 有几个 Key 动


### 基本概念

一、Register：服务注册
它提供自当 Eureka客户端向 Eureka Server注册时，
身的元数据，比如IP地址、端口，运行状况指示符URL，主页等。

二、Renew：服务续约
Eureka客户会每隔30秒发送一次心跳来续约。 
通过续约来告知 Eureka Server 该Eureka客户仍然存在，没有出现问题。 
正常情况下，如果 Eureka Server 在90秒没有收到 Eureka 客户的续约，它会将实例从其注册表中删除。 
建议不要更改续约间隔。

三、Fetch Registries：获取注册列表信息
Eureka客户端从服务器获取注册表信息，并将其缓存在本地。
客户端会使用该信息查找其他服务，从而进行远程调用。
该注册列表信息定期（每30秒钟）更新一次(d代码逻辑）。
每次返回注册列表信息可能与Eureka客户端的缓存信息不同， 
Eureka客户端自动处理。如果由于某种原因导致注册列表信息不能及时匹配，
Eureka客户端则会重新获取整个注册表信息。 
Eureka服务器缓存注册列表信息，整个注册表以及每个应用程序的信息进行了压缩，压缩内容和没有压缩的内容完全相同。
Eureka客户端和 Eureka 服务器可以使用JSON / XML格式进行通讯。
在默认的情况下 Eureka客户端使用压缩JSON格式来获取注册列表的信息。

四、Cancel：服务下线
Eureka客户端在程序关闭时向Eureka服务器发送取消请求。 
发送请求后，该客户端实例信息将从服务器的实例注册表中删除。该下线请求不会自动完成，它需要调用以下内容：
DiscoveryManager.getInstance().shutdownComponent()；

五、Eviction 服务剔除
在默认的情况下，当Eureka客户端连续90秒没有向Eureka服务器发送服务续约，即心跳，
Eureka服务器会将该服务实例从服务注册列表删除，即服务剔除


### 什么是微服务？

1. 什么是服务发现？
2. 服务发现应该具备哪些关键特性，理由是什么？
3. 你认为服务发现带来的主要好处是什么？
4. 时至今日，哪一种服务发现方案是最可靠的？
5. 实施服务发现面临的最大挑战是什么？
6. 在一个全新的系统中实现服务发现（相对）容易。已有的系统如何集成服务发现功能？


### eureka 微服务
1. https://github.com/Netflix/eureka/wiki

### Eureka 服务
1、And further more, Eureka has been built carefully 
without any hard dependency on any external components .

### 如何在 Discovery 应用负载均衡算法？

### 如何才能检测到 Data Race ? 
1. [例子如下图所示] https://github.com/bilibili/discovery/pull/78
2. 


### 为什么需要服务发现？
1. https://zhuanlan.zhihu.com/p/161277955
一、1. 什么是服务注册与发现
二、web1.0数据请求模型架构, 2.0, 3.0 的相关架构（包括和现在的架构是否一样？）
三、微服务时代的服务管理

2. FAQ
为什么有了 k8s 还需要服务发现？
利用 k8s 内部的负载， 用 service_name + port 就可以定位到某个实例子了。
 
 
### 为什么要使用 Euraka 作为服务发现？
1. https://medium.com/knerd/eureka-why-you-shouldnt-use-zookeeper-for-service-discovery-4932c5c7e764
2. 