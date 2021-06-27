
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


### 疑惑 
1、k8s yaml 中的 image 是如何被拉去的？ 机制是什么？
各个工作节点负责运行 Pod 对象，而 Pod 的核心功能在于运行容器，因此工作节点上必须配置容器运行引擎。
使用私有仓库中的惊喜时，通常需要由 registry 服务器完成认证后才能够进行。 

2、暴露端口，如何查看 k8s 集群上/node 上暴露了哪些端口？

3、环境变量， env 和 envFrom 

4、共享节点的网络和名称空间。

spec.hostNetwork 的属性为 true, 代表什么？
kubectl exec -it pod-use-hostnewowork — sh ， 验证一下工作节点的网络设备及相关的接口信息。 

5、标签与标签选择器 Label Selector。 
常用的标签， 版本标签、环境标签、应用标签、架构层级标签、分区标签、品控级别标签
《Kubernetes in action 》

6、标签选择器， 有哪两种？有什么用？
metadata 中嵌套使用 labels 字段。 
标签赋予 Kubernetes 灵活操作资源对象的能力， 它也是 Service 和 DeployMent 等核心资源得以实现的基本前提。 

7、Pod 节点选择器 nodeSelector 

8、标签注解， metadata.annotations 或者通过 kubectl annotate 添加。 

9、POD 对象的生命周期。 有几种状态（phase)？ https://blog.openshift.com 

10、Pod 的启动过程。***（难点） 

11、Pod 生命周期中的重要行为 ***（难点） 

12、kubectl get pods -l test=xxx -exec -w 
运行 “ kubectl get -w “ 命令监视其资源变动信息。 

13、 设置 TCP 探针、HTTP 探针
readinessProbe、livenessProbe 之间有什么区别？

14、容器的可见资源（request 和 limit 有什么区别） ？
例子， 在 POD 运行 nginx 应用， 在配置参数 worker_processes 为 auto 时，注进程会创建与 pod 能够访问到 CPU 核心数相同数量的 worker 进程。 这有可能带来问题， 需要利用 Downward API 进行解决。 


15、存活性探测、就绪性探测是辅助判断容器状态的重要工具🔧。

16、Pod 控制器，对 Pod 对象的管理通常是由某种控制器的特定对象来实现的，包括其创建、删除以及重新调度等操作。 
三个级别组成部分， 标签选择器、期望的副本数、Pod 模版。

17、loop-reconciliation 和 list-watch 是 k8s 核心机制。 在资源对象状态发生变动时，由 API Server 负责写入 etcd 并通过水平触发 ( level-triggered ) 机制主动通知给相关的客户端程序以确保其不会错过任何一个事件。 

18、Deployment 控制器，事件和状态、回滚、版本记录、暂停和启动、多种自动更新方案。 
Deployment 借助与 ReplicaSet 管理 Pod 资源的机制。 

