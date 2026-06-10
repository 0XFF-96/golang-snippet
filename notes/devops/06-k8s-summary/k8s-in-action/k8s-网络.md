### k8s 网络模型

docker 容器网络主要由三种： Bridge （桥接）、Host (主机）、Container (容器） (多个容器共享一个网络名称空间） 。 这些都是解决 在同一节点上的容器通讯的可用方案。 跨节点和跨容器的，在生产环境更为常见。 k8s 网络模型

1、容器间通信， pod 内的容器间通信，是用过 lo 接口完成交互。 如同同一主机上的多个进程类型。 

2、pod 间通信， 每个 Pod 对象拥有一个集群全局唯一的地址并可用直接用于与其他 Pod 进行通信。 类于同一 IP 网络株机间进行的通信。 这是重点着力解决的问题。 

3、Service 与 Pod 间的通信， 重点在与 Service 的 Cluster-IP 和 Pod-IP 之间的报文转发，及其背后的实现。

4、集群外部到 Pod 对象之间的通信， 前提条件是 Pod 所在的工作节点范围和节点端口 ( nodePort ) 和主机网络 ( hostNetwork ) 两种， 以及工作于集群级别的 NodePort 或 LoadBalancer 类型的 Service 对象。CNI 插件及常见的实现
1、k8s 设计了网了模型，但其实现交给了网了插件。 CoreOs 和 Google 联合制定了 CNI 标准， 定义了容器网络模型规范。 flannel 网络插件， 依赖于 VxLan 
1、各种 CNI 插件需要解决的两类问题是， 一、不同节点的容器可能得到相同的地址，造成地址冲突问题。 
2、多个节点的 docker0 桥使用不同子网， 其报文因为在网络缺少路由信息，而无法准确到达。 
3、flannel 使用 etcd 存储虚拟 IP 和主机 IP 之间的映射，各个节点运行 flanneid 守护进程负责监视 etcd 中的信息并完成报文路由。 4、配置在 /http://coreos.com/network/config  etcd 键值下。什么是 Vxlan 1、虚拟可以扩展局域网， 采用的是 MAC in UDP 封装方式。 配置 k8s 的网络策略1、同一 namespace 和不同 namespace 之间的网络隔离。 

### 网络策略应用案例 p534 

1、同一个 pod 内的多个容器之间：lo 
2、各个 Pod 之间的通讯： Overlay NetWork。 CNI 接口， Flannel 设计的一个网络规划服务。 
3、Pod 与 Service 之间的通讯， 各个节点的 Iptables 规则。


### 共享节点（Node) 的网络名称和空间

1、Pod 对象的容器，均运行与一个独立的、隔离的 NetWork 名称空间中，共享同一个网络协议栈及相关的网络设备。 

2、当 开启呢 spec.hostNetWork 的属性为 true 时，即可创建共享节点网络名称空间的 Pod 对象。 此类 pod 为， kube-scheduler、kube-proxy、kube-flannel 等。 