### helm 的使用
1、https://www.mls-tech.info/microservice/k8s/minikube-use-helm/

2、helm 的服务端安装，以及和 k8s 之间的通讯。 
因为 Helm 的服务端 Tiller 是一个部署在 Kubernetes 中 Kube-System Namespace 下 的 Deployment，它会去连接 Kube-Api 在 Kubernetes 里创建和删除应用。 

4、https://www.hi-linux.com/posts/21466.html

5、https://zhuanlan.zhihu.com/p/69274854
book : https://jimmysong.io/kubernetes-handbook/

6、k0s 项目， https://github.com/k0sproject/k0s

### helm 基础概念

1、helm 是应用程序包管理器， 类似于 yum 和 apt-get ， 核心打包功能组件为 chart 
2、helm 将 k8s 资源， 打包到一个 Charts 中，制作并测试完成的各个 Charts 将报错到 Charts 仓库进行存储和分发。 
3、核心概念： Charts、Repository、Config、Release
4、架构 Helm 客户端、Tiller 服务器和 Chats 仓库组成。 常用命令：
1、helm list 
2、helm delete xxx 
3、helm install 
4、helm search 
5、helm inspect 
6、helm repo update Charts 及其概念。 


### helm 的部署和安装 


1、tiller 如何交互？如何按照 tiller 服务端
2、我之前按照过 helm, 但是利用 minikube 删除过 k8s 集群好几次了，为什么我的 helm 客户端，还能依旧成功成功连接到我的集群？ kube-config ？
3、https://github.com/helm/charts 


### Ref 
1. https://www.mls-tech.info/microservice/k8s/minikube-use-helm/
