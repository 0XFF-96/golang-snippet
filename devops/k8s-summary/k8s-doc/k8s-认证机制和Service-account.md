### k8s 准人和认证机制

- 你为什么能够操作 kubectl ?
- pod 之间访问 k8s api 如何交互？


1、理解 人 和 “POD 对象“ ， 用于区分 User Account 和 服务账号 Service Account 
2、什么是 k8s 中的服务账号？ 
服务账户：其使用主体是“应用程序”，专用于为Pod资源中的服务进程提供用于访问 Kubernetes API 时的身份标识（identity）；ServiceAccount资源通常要绑定于特定的名称空间，它们由API Server自动创建或通过API调用由管理员手动创建，通常附带着一组存储为Secret的用于访问API server的认证凭据，可由同一名称的Pod应用访问API Server时使用

3、Kubernetes借助于同名的 ServiceAccount准入控制器（Admission Controller）、Token 控制器 和 ServiceAccount 控制器三个组件实现了 ServiceAccount 的自动化。

4、获取集群的 kubeconfig 文件，并写入当前用户目录下。 这个相当关键。 如何通过配置 kubectl , 访问到阿里云上面的 k8s 就是靠这个。 

Secret 对象的特色类型。 
ca.crt      # CA的证书文件
namespace   # 保存有适用的名称空间的文件
token       # 保存有认证到API Server的Token的文件

ServiceAccount是标准的资源类型，但它不支持使用spec字段，而是直接指定对象名称后，由Kubernetes系统自动完成后续的功能，例如创建关联的Secret及Token等。 

ServiceAccount资源还可以基于 spec.imagePullSecret 字段附带一个由下载镜像专用的 Secret 资源组成的列表，用于让Pod对象在创建容器时从私有镜像仓库下载镜像文件之前完成身份认证。事实上，我们实验环境中所使用的ACK集群上自动生成的serviceaccounts/default就被自动附加了几个 imagePullSecret。 

以curl命令向API Server发起访问请求，我们通过附加承载令牌的方式模拟它以该Pod上关联的ServiceAccount的身份请求列出当前名称空间下的Pod对象

6、所有 Service Account 的主要内容是用于与 API Server 交互的对吗？


1、为什么你能够使用 kubectl ?

2、https://kubernetes.io/docs/reference/access-authn-authz/authentication/#service-account-tokens

—kube-config , 这个用于连接 k8s 的命令的相关密码，来源于哪里？



### 准入、认证、授权


1、k8s 怎么拉到某些镜像？
2、怎么确定这个用户是否有权限操作此资源对象？
3、用户账号、服务账号、认证、授权和准入控制。 

API-Server 支持哪几种具体的认证方式？

API-Server 内建的授权插件有哪些？
1、Node, 2、ABAC: 基于属性的访问控制
3、RBAC: role-based access control 基于角色的访问控制。 
服务账户， 让 Pod 对象内的容器进程访问其他服务时，提供身份认证信息的账户。 Service Account 资源一般由用户名以及相关联的 Secret 对象组成。 

pod 对象与服务账号，
创建 Pod 对象时，可为 Pod 对象指定使用自定义服务账号，从而实现自主控制 Pod 对象资源的访问授权权限。 

pod 如何拉去镜像？
ServiceAccount 资源可以基于 spec.imagePullSecret 字段，带下一个由下载镜像专用的 Secret 资源组成的列表，用于在进行容器创建时，从某私有镜像长裤下载镜像文件之前进行服务认证。 
kubectl config view , 
使用 kubeconfig 配置，介入多个集群的相关配置信息， 并且可以在各个环境之间切换。 二、Role 和 RoleBinding, Role 仅时一组许可 （permission ) 权限的集合，它描述了对哪些资源可执行何种操作，资源配置清单中使用 rules 字段嵌套授权规则。 

1、为什么你能够使用 kubectl ?
2、https://kubernetes.io/docs/reference/access-authn-authz/authentication/#service-account-tokens
—kube-config , 这个用于连接 k8s 的命令的相关密码，来源于哪里？


