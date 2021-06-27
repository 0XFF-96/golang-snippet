### 什么是 client-go 
1、可以说 Kubernetes 是当前云原生的基石。所以想要进军云原生领域，kubernetes 的学习必不可少。kubernetes 的设计理念就是通过各种控制器将系统的实际运行状态协调到声明 API 中的期待状态。而这种协调机制就是基于 client-go 实现的。同样，kubernetes 对于 ETCD 存储的缓存处理也使用到了 client-go 中的 Reflector 机制。所以学好 client-go，等于迈入了 Kubernetes 的大门

REF:
1、https://github.com/opsnull/kubernetes-dev-docs/tree/master/client-go
2、Get Started, https://minikube.sigs.k8s.io/docs/handbook/controls/ 。 