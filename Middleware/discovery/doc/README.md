### 服务注册与发现学习

### 什么是服务发现？
1. http://dockone.io/article/509 [done, 什么是微服务]

1. AP( Availability and Partition ) 系统， Eureka 是一个很好的选择，
并在 Netflix 得到了实战的检验。在出现网络分区时， Eureka 选择可用性，而不是一致性

### 为什么需要服务发现？
1. https://toutiao.io/posts/9c7awe/preview

一、为什么需要服务发现？（客户端和服务端的划分机制） 

服务实例会被动态地分配网络地址。并且，因为自动伸缩、故障和升级，服务实例会动态地改变。
故而，你的客户端代码需要用一种更加精密的服务发现机制。
有两个主要的服务发现模式：客户端发现和服务端发现。让我们先来看看服务端发现

### 服务发现有哪几种方式？
1. https://xie.infoq.cn/article/13a6973621381b27bbcd9ab45 
 
观点1： 对于服务发现与注册 一般有两种实现模式：服务器端模式，客户端模式 
（为什么分两种方式) 
   

### 知识点
1. 在 Nodes 中，应该如何使用 error group
2. 什么是自我保护机制？https://mp.weixin.qq.com/s/fr5m5x1wv8zQ4W2_po_9_g

### 参考 
1. https://zhuanlan.zhihu.com/p/88266860 
2. eureka 之间的区别，https://zhuanlan.zhihu.com/p/84659866 
3. idea from eureka https://blog.csdn.net/forezp/article/details/73017664 
4. eureka 服务发现剖析，https://xie.infoq.cn/article/c6139a5776e139d1b76b2e856
5.  