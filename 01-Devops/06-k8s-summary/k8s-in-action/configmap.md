
1、了解硬编码环境变量的不足，以及不能复用
2、创建 configmap 的方式，从文件内容、文件夹创建 configmap 
3、在 pod 中引用不存在的 ConfigMap, 容器启动会失败。 除非设置 configMapKeyRef optional = true 
4、一次性传递 ConfigMpa 的所有条目作为环境变量。 
5、传递 ConfigMap 条目作为，启动 Pod 时的命令行参数。 
6、使用 configMap 卷将条目暴露为文件。 
7、更新应用配置且不重启应用程序， 利用 ConfigMap 达到热更新的目的。 
Secret 存储 Ingress 资源的 TLS 证书。 

1、default-token 的作用， 包含三个条目， ca.crt、namespace 与 token 。 包含了从 pod 内部安全访问 k8s API 服务器所需要的全部信息。 
这个 secret 被默认挂载到每个容器， 可以通过 automountServiceAccountToken 字段设置为 false . 
2、dig 命令， nslook up 命令。 https://www.cnblogs.com/zgq123456/p/9935597.html 。 
3、secret 大小限制与 1 MB , 被 base64 编码。
4、利用环境变量。 
5、如何通过加 secret 到 service account , 使所有 pod 都能自动添加上镜像 secret 。 