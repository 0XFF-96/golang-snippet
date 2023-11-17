### Downward-api 

1、Downward API 可以给在 pod 中运行的进行，暴露 pod 的元数据。 有以下。 
2、通过环境变量暴露元数据。
3、通过 downwardAPI 卷来传递元数据。
4、通过与 k8s API 服务器交互，获取集群元数据。
5、通过 kubectl proxy 访问 API 服务器， 用于研究 Kubernetes API 。 
6、在 minikube 中关闭 RBAC 机制，服务账号不受限制地获取任何信息。 