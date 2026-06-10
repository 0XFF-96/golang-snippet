1、为何多个容器壁丹哥容器包含多个进程好？单进场容器， 方便管理和监控。 
2、同一个 Pod 中容器之间的部分隔离
3、容器如何共享相同的 IP 和 端口空间？
4、pod 的设计， 将多层应用分散到多个 pod 中， 基于扩容考虑而分割多个 pod 
5、在 pod 中使用多个容器的条件。 
–	他们需要一起运行，还是可以运行在不同主机上
–	他们代表一个整体，还是独立的组件
–	他们可以一起进行扩缩容吗
6、可以利用 kubectl explain pods ， 进行调试。查看每个 API 对象支持哪些属性。 
7、每次/日志文件到达 10 MB ,容器日志都会进行自动轮替， kubectl logs 命令仅显示最后一次轮替后的日志条目。 
8、通过 port-forward ，将本地网络端口转发到 Pod  中的端口。 
9、使用标签组织 Pod,  k-v 的方式。可以查看常用的标签。10、通过标签选择器列出 pod 子集。 
–	kubectl get po -l env 
–	kubectl get po -l ‘!env’ 
11、使用标签和选择器，约束 Pod 调度、分类工作节点、调度到特定的 Pod。
12、注解和标签的区别？
13、kubectl get ns , 列出命名空间
14、命名空间的隔离， 命名空间是否提供网络隔离取决于 Kubernetes 所使用的网络解决方案。 
15、使用标签删除 pod ， kubectl delete po -l ‘label_name’ 一、更改 RelicationController 的标签选择器
1、会导致之前管理过大 Pod, 脱离 RC 的控制，变成自主式 pod 。 
2、使用 ReplicaSet, 而不是 ReplicationController 
3、RS 运行匹配缺少某标签的 Pod , 
4、RS 比 RC , 具有更表达力的标签选择器。 matchExpressions 。 
5、kubectl get pods -A 
6、在 Job 中运行多个 Pod 实例。 
7、配置 k8s 的 cronjob 模式。