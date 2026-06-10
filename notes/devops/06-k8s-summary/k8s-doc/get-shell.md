1、https://kubernetes.io/docs/tasks/debug-application-cluster/get-shell-running-container/


2、Gettting into containers via exec 、proxies、port forwarding . 

3、what is exec, 


一、hands-on example 
1、kubectl apply -f https://k8s.io/examples/application/shell-demo.yaml

2、kubectl get pod shell-demo

3、kubectl exec --stdin --tty shell-demo -- /bin/bash

4、kl get pods shell-demo -o yaml
遇到端口占用的错误， 80 端口。 这是绑定在了 node 的 80 端口上？
nginx: [emerg] bind() to 0.0.0.0:80 failed (98: Address already in use), 看起来是因为这个， 使用hostNetwork参数（容器内部服务与宿主机同一网段）。 当Pod调度到哪个节点就使用哪个节点的IP地址，客户端使用IP地址访问容器里面的服务。一个node只能启动一个pod端口，端口不能冲突。 删除了另一个 nginx pod 就解决了， 但是这个 pod 没有设置 hostNetWork = true, 的参数， 只用了 containerPort = 80 , 为什么会造成冲突？
app-service                      NodePort    10.111.244.51    <none>        80:30080/TCP         41h这里的 80 -> 30080 , 是什么意思？ 从哪里映射到哪里？最有一种调试方式， 开一个零食的 pod, 暴露和 host / node 一样的网络， 然后在上面执行 netstat 命令， 看看哪个应用占用了端口。 这种方法执行 netstat, 是否能够查看 node 节点的相关信息？不一定的吧。 
https://www.jianshu.com/p/59b62e024072 ， lsof -i 【端口号】https://cloud.tencent.com/developer/article/1637116 
5、如果查看 node 的 80 端口，有没有被占用？
6、kubectl  get pods -o wide
7、没有找到具体 80 port 被占用的原因。 修改了 hostNewWork 参数之后， pod 运行正常。 
8、kubectl exec --stdin --tty shell-demo -- /bin/bash
9、ps aux | grep nginx ， 这个是什么命令？单独运行命令：
kubectl exec shell-demo -- ps aux
kubectl exec shell-demo -- ls /
kubectl exec shell-demo -- cat /proc/1/mounts
当 pod 内多于一个容器时候， 
 kubectl exec -i -t my-pod --container main-app -- /bin/bash