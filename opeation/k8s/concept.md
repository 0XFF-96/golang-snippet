
### 阿里云 （知行动手实验室的笔记）
1、什么是 k8s 上的对象？
Kubernetes的API 元件（Primitive）用于描述在Kubernetes上运行应用程序的基本组件，这些元件也就是俗称的Kubernetes对象（Object）
2、spec、status 字段的意义？
3、kind 和 apiVersion 两个字段的含义？
4、k8s 有哪些核心资源类型？Kubernetes的核心资源类型
Kubernetes的API对象大体可分为工作负载（Workload）、发现和负载均衡（Discovery & LB）、配置和存储（Config & Storage）、集群（Cluster）和元数据（Metadata）几个类别5、pod 有哪几种分类？6、kubectl 命令提供了三种类型的对象管理机制。
```
kubectl get pods -o wide
kubectl get pods --all-namespaces
```
6、查看 etcd 原信息的相关方法？
podName=$(kubectl get pods -o "jsonpath={.items[0].metadata.name}");kubectl get pods $podName -o yaml --show-managed-fields=false
我们先获取任意一个Pod对象的名称后，打印其完整格式的资源规范
7、在 pod 内的容器，运行相关的命令。
kubectl exec pods/pod-demo -c demo -it -- netstat -tnlp
7、如何定制要运行的容器应用及参数？ 
8、在Pod上使用环境变量进行应用配置？ 有几种方式 


### From k8s 进阶实战

1、kubernetes 集群中的每个节点都运行着 cAdvisor 以收集容器及节点的 CPU、内存及磁盘资源的利用率指标数据，这些统计数据由 Heapster 聚合后通过 API Server 访问。 
2、kubectl exec , 在 pod 对象某容器内运行指定的程序，其功能类似于 “ docker exec” 命令，可以用于了解容器各方面相关信息或执行必须的设定操作等，其具体功能取决于容器内可用的程序。 
3、进入某 pod 内的容器，kubectl -it exec $POD_NAME /bin/sh 
4、k8s 有三个核心资源抽象 Pod、Deployment、Service ，并介绍了 kubectl 的基本用法之后，通过案例讲解了如何在集群中部署、暴露、访问及扩容容器化。 
5、k8s 的资源对象 和 用户自定义资源。 
自定义资源 （ Custom Resource Definition CRD ) 
6、  kubectl proxy ？ 这个命令有什么用？
P113 页面。
7、什么是 namespace ?
Namespace 对象仅用于资源对象名称的隔离，自身并不能隔跨名称空间的 Pod 间通讯，那是网络策略(Network policy) 
8、暂停 （pause ) 或者 继续 （resume) 更新操作。maxSurge 和 maxUnavailable 属性还能实现更为精巧的过程控制。 
9、DaemonSet 控制器，特殊的应用场景。通常运行执行系统级操作任务的用用。运行集群存储的守护进程、运行日志手机守护程序 fluentd ，在各个节点上运行监控系统的代理守护进程。


### 什么是 Spec 和 status 字段？ 

- Spec,  用于描述所期望的对象应该具有的状态，必须具有的。 
- Status, 用来记录对象在系统上的当前状态， 此字段由 k8s 负责填充，不能由用户更改。 Master 的 controller-manager 通过相应的控制器组件动态管理并确保对象的实际状态，匹配用户期望的状态，它是一种调和 ( reconciliation ) 配置系统。 
kubectl explain pods.spec 
