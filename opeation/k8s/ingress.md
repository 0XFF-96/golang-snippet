### 部署 Ingress 服务

命令： 1、kubectl get pods -n kube-system nginx-ingress-controller-6fc5bcc8c9-khm5p -o yaml
2、部署一个小型 go 应用。

执行 
- kl get svc go-app-svc -o yaml  检查 yaml  
- minikube ip , 和 kl get svc 当中的映射地址， 进行访问，看是否能够从本地网络访问出去。 

3、准备进行扩容
- kubectl rollout status deployment my-go-app
- kubectl get replicaset
- kubectl scale --replicas=3 deployment my-go-app --record
- kubectl describe deployment my-go-app
- kubectl rollout history deployment my-go-app 【历史 rollout 记录】 
- kubectl rollout undo  deployment my-go-app --to-revision=1 【回滚 deployment 】

如果中途遇到一些错误，get event 可以知道，是节点的 CPU 不够。 Waiting for deployment "my-go-app" rollout to finish: 2 of 3 updated replicas are available...

4、控制 ReplicaSet 的版本数量
–	kubectl rollout pause deployment my-go-app 
5、部署 service.yaml 文件 ， 
–	kubectl apply -f service.yaml 。 在 class/k8s/ingress 目录中。
–	kubectl get ep app-service 
–	spec.ports配置里 ports : nodePort, port, protocol, targetPort , 这几个名称背后代表的含义是什么？
–	kl exec -it my-go-app-5f87c4cb78-6jd9h -- /bin/sh ， 进入 service 所在的 pod 里面。 
–	nslookup app-service.default.svc.cluster.local , 【查看这条 DNS 记录】 【需要在所在 Pod 中执行】
–	nslookup *.app-service.default.svc.cluster.local  【和上面一条查询的内容为什么会有不同？】
6、部署 app-ingress 文件, 所有对 http://app.example.com 入口的请求都路由到app-service这个Service的80端口。 
–	kubectl apply -f app-ingress.yaml --record
–	kubctl describe ingress
–	在域名中增加， 192.168.64.6 （minikube ip ) http://app.example.com 
–	curl [app.example.com ,](http://app.example.com)    配置完成后，能够看见与上面一模一样的结果输出。 


### 几个 Port 的含义 
- port, 指集群内部暴露 Service 所使用的端口， 集群内部使用 <ClusterIP>:<port> 访问 Service 的 EndPoints （ Service 所选中的 Pod )

- nodePort, 指向集群外部暴露 Service 所使用的端口， <NodeIP>:<NodePort>, 访问 Service 的 Endport 

- targetPort, 是后面 Pod 监听的端口，容器里面的应用也应该监听这个端口，Service 会把请求发送到这个端口里面。 怎么发现 Service ?
在 kubernetes 里的内部组件 kube-dns 会监控 kubernetes API, 当有新的 Service 对象被创建出来后，kube-dns 会为 Service 对象添加 DNS A 记录 （ 从域名解析 IP 的记录） ， 对于 【 ClusterIP 】模式的Service ，它的 A 记录格式是： serviceName.namespace.svc.cluster.local , 当访问这条 A 记录时，解析到的就是 Service 的 VIP 地址。 


### Ref 

1. http://mp.weixin.qq.com/s#wechat_redirect?__biz=MzUzNTY5MzU2MA==&mid=2247486082&idx=1&sn=42a9bc8fcfc9da09445e9e2f4cf2fb96&chksm=fa80d%E2%80%A689ee&token=1964476830&lang=zh_CN&scene=21 
3、http://mp.weixin.qq.com/s#rd?__biz=MzUzNTY5MzU2MA==&mid=2247486105&idx=1&sn=4d99b13a2a67a70d506f446a7fad10a9&chksm=fa80d%E2%80%A60a6c&cur_album_id=1394839706508148737&scene=189 

2. https://kubernetes.io/zh/docs/concepts/services-networking/ingress/ 


3. 部署一个小型应用。 https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247485328&idx=1&sn=1615cbb8d60b0df38cf59375cca6cef2&chksm=fa80d607cdf75f11cf025730a492646f3ea375a81fca5e70c46703f37aa2e7d929950693cae6&token=2051310148&lang=zh_CN&scene=21#wechat_redirect

