### 让客户端发现 pod 并通信

> 1、https://kubernetes.io/zh/docs/concepts/services-networking/service/ 


1、sessionAffinity, ClientIIP . 这种方式会使服务代理将来自同一个 client ip 的所有请求转发至同一个 pod 。 
2、命名端口的好处是， 即使更换端口好，也无须更改服务 spec 。 
3、环境变量是获得服务 IP 地址和端口号的一种方式。 
4、DNS 发现服务， k8s 会修改每个容器的 /etc/resolv.conf 文件实现， 运行在 pod 上的进程 DNS 查询都会被 k8s 自身的 DNS 服务器响应。 
5、Endpoint 的名称必须和服务的名称相匹配。 
6、将服务暴露给外部客户端。 NodePort, 每个集群节点上打开一个端口， 并且在改端口上接受到的流量重定向到基础服务。 
7、将服务的类型设置为 LoadBalance, NodePort 的一个扩展， 是的服务可以通过一个专用的负载均衡器访问。? 这个不懂 。 这个和 ExTERNAL-IP 有关。 
8、创建一个 Ingress 资源，反向代理。只需要一个公网 IP， 就能为许多服务提供访问。
–	将不同服务映射到相同主机的不同路径。
–	将不同的服务映射到不同的主机上。 
9、externaltrafficPolicy: local 配置的优缺点。 10、ingress 控制器，负载处理与 TLS相关的内容， 需要将证书和私钥附加到 Ingress 上。 可以保持在 Secret 资源中。 

11、使用 headless 服务来发现 【独立 pod 】。 如果客户端需要连接到所有的 pod ? ,  k8s api 服务器通过 api 获取 pod 及其 ip 地址列表。 k8s 通过 DNS 查找发现 pod IP, 为该服务返回多个 A记录。 客户端可以使用信息连接到其中的一个、多个或者全部。 
12、集群 IP 是虚拟 IP， 是无法 ping 通的。
13、为具有 ExternalName 服务类型的外部服务提供 DNS CNAME 别名。  