### k8s 日志架构

一、集群级日志
在集群中，日志应该具有独立的存储和生命周期，与节点、Pod 或容器的生命周期相独立。 
2、需要一个独立的后端用于存储、分析和查询日志。
3、logs 文档， https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#logs 。 
4、节点级日志架构和集群级日志架构之前的关系。
5、节点级日志记录（图，需要手动画出来） logrotate、log-file.log 和 stdout、stderr 之间的关系。 

6、如果容器重启、Pod 索价节点被驱逐， 会发生什么？

7、节点上的日记记录需要考虑如何实现日志的轮转，保证不会消耗节点上的可用空间。 
8、kubectl logs 的工作原理。

二、系统组件日志
1、分为两类： 在容器中运行 和 不在容器中运行
2、开发文档，https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md
3、/var/log 目前中的系统组件日志，每天/超过 100 MB 时触发轮转。

三、集群级日志架构方案
1、我们现在采用的是哪种方案？

### 容器日志

1、Pod中各容器的文件系统彼此隔离，且与宿主机的文件系统也互相隔离，使用了宿主机上的文件系统作为存储卷的场景除外。于是，容器中的日志信息若存储于容器内部的文件系统上，将会导致收集的困难，而且也会随着Pod的删除而丢失这些日志信息
2、从零，搭建一个 k8s 或者是 staging 集群需要多久？需要一些什么资源和前提？
3、kubectl logs $podName -c proxy ， 这条命令的具体操作手法是什么呢？
4、kubectl logs $podName --since=5m -c demoapp
5、打印上一个 pod 的日志。 

二、构建一个集群日志系统。 
1、ElasticSearch、Fluentd（或Fluent-bit/Filebeat等）和Kibana
2、PLG：Promtail、Loki和Grafana
https://kubernetes.io/zh/docs/concepts/cluster-administration/logging/?spm=5176.12026607.tutorial.2.9c7732e2CrsyJf
3、kubectl get ingress
4、https://www.qikqiak.com/post/visually-explained-k8s-ingress/
5、部署日志收集代理（fluent-bit）
持ClusterRolebinding等资源类型上的管理操作
6、kubectl get statefulsets,deployments,services,ingress,pvc

### 资源
1. https://kubernetes.io/zh/docs/concepts/cluster-administration/logging/#sidecar-container-with-a-logging-agent


