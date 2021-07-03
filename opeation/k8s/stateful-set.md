### stateful-set 相关
大纲
1. 用 k8s 编排有容器应用
2. sattefulset 实践 
3. 什么是 【拓扑状态】 和 【存储状态】

一、用 kubernetes 编排有状态应用。
 
VolumeClaimTemplate， 来声明 Pod 存储用的 PVC, 当部署 StatefulSet 部署 Pod 时，会从编号 0 部署到最终个 Pod 

2、如何维护 【拓扑状态】 和 【存储状态 】？

ClusterIP 模式的 Service 和 Headless Service 的区别？Service是在逻辑抽象层上定义了一组Pod，为他们提供一个统一的固定IP和访问这组Pod的负载均衡策略

- A 记录时一样的， serviceName.namespace.svc.cluster.local

- headless service 访问记录后返回的 IP 地址集合， Pod 也会被分配对应的 DNS A 记录， 如， podName.serviceName.namesapce.svc.cluster.localHeadlessService 。 

3、StatefulSet 需要 Headless Service 的原因是， 它会为代理每个 StatefulSet 创建出来的 Endpoint ， Pod 添加 DNS 域名解析，Pod 之间就可以相互访问。 2、StatefulSet 给管理的所有 Pod 名字，进行了编号。 StatefulSet 名号-序号， 从零开始累加。  -  kl create -f headless-service.yaml 
–	进入所在 pod 查看，相关的 DNS 记录。 
–	kubectl exec -it my-go-app-69d6844c5c-gkb6z  -- /bin/sh
–	kubectl exec -it stat-go-app-0  -- /bin/sh 
–	nslookup app-headless-svc.default.svc.cluster.local  【查看 dns 记录】
–	nslookup stat-go-app-0.app-headless-svc.default.svc.cluster.local 【查看 pod 】的 dns 记录。 

二、statefulset 的几个作用
1、保持 Pod 的编排顺序， 
2、保持 Pod 固定唯一网络标识， Pod 之间互相能以 podName.serviceName.namesapce.svc.cluster.local  ， 通信，不担心 Pod 被重新调度到其他节点后的 IP 变化。 
3、保持实的存储状态， spec.volumns, host path 类型的 Volumn 基于宿主机目录的， Pod 发生重新调度后，就不能恢复了。 所以，此功能需要使用集群存储资源， PV 和 PVC 。 分布式存储项目 （Ceph、GFS、HDFS 等） 


### Ref

1. https://mp.weixin.qq.com/s/0rSutWfY_ENpvIidbp8ePw网管叨bi叨

