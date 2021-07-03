### 什么是 k8s dashboard 
1、kubectl get svc kubernetes-dashboard -n kube-system 
2、配置 kubeconfig 认证， 用于登录 k8s-dashboard 
如何配置一个仅具有特定名称空间管理权限的登录账号
- 首先新建一个 ServiceAccount ， 用于管理默认的 default 名称空间。 kubectl create serviceaccount def-ns-admin -n default 
- 将之绑定与 admin 集群角色。 
- 创建所需的 kubeconfig 文件，初始化集群、提供 API Server、验证 API Server 证书所用到的 CA 证书。 

### 总结一下
1、k8s Dashboard 及其分级权限的授予。
2、Service Account 资源及其应用方式。
3、Role 及 RoleBinding 的工作机制及应用方式。
4、ClusterRole 及 ClusterRoleBinding 的工作机制及应用方式。