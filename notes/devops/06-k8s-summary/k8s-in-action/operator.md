### Operator 模式

1、什么是 operator 和如何创建 operator 

2、Kubernetes Operator要做的：代替原本需要由SRE（Site Reliability Engineers）和运维工程师来完成操作的执行。

3、Operator 模型， 自定义资源和自定义控制器 。 

4、什么是 reconcile ? 

5、kubernetes 的所有控制器， 都有一个控制循环，负责监控集群中特定资源的更改，并确保特定资源在集群里的当前状态与控制器自身定义的期望状态保持一致

6、Operator 都是控制器，但是并非所有控制器都是 Operator 