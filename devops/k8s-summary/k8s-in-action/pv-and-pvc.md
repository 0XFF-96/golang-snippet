### 什么是 PV 和 PVC 

1、PVC 和 PV 的关系相当于， 编程领域的 接口和实现。 persistentVolumeClaim、persistentVolume

2、PVC 表达的是 Pod 的存储要求， 有一种解耦合的思想在里面。 我可以在完全不知道远程存储的、服务器地址、AccessKey 之类的信息。 

3、StatefulSet 额外添加一个 spec.volumnClainTemplates 的东西。

4、StatefulSet 创建的 PVC, 都以 PVC名-Stateful名-序号这个格式命名的。 5、StatefulSet 定义文件中的 volumnClaimTemplates 声明 Pod 使用的 PVC， 它建出来的 PVC 会以名称编号这些约定与它创建出来的 Pod 进行绑定，借由 PVC 独立于 Pod 的生命周期和两者之间的绑定机制的帮助，StatefulSet 完成了应用存储状态的维护。 

### 存储类 （二） 
1、对于存储资源的使用最好也能像使用计算资源一样，用户和开发人员无限了解 Pod 资源究竟运行于哪个节点，也无需了解存储系统是什么设备以及存储何处 。 

2、PVC 和 PV 是一对一的关系。 PV 是集群级别的资源，而 PVC 属于命名空间。

3、存储类， storage class 。 由管理员为管理 PV 而按需创建的类别 （ 逻辑组）。 可按存储系统的性能高低分类， 或者根据其综合服务质量级别进行分类。 

4、还有一些概念不太懂， PV 和 PVC 的生命周期。 存储供给、存储绑定、存储回收。 

### 存储卷

存储卷的基本使用

1、存储卷类型， emptyDir, hostPath （节点基本的卷类型）  还有一些网络相关的类型。 
2、PV PersistenVolumn 是一种集群级别的资源。 
3、CSI 容器存储接口，1.9 开始引入。
4、spec.volumns、spec.containers.volumnMounts 字段 有什么区别？

5、gitRepo 存储卷， emptyDir 存储卷的一个实际应用，可以通过挂载目录访问指定的代码仓库中的数据。

6、hostPath, 指将工作节点上某文件系统的目录或者文件挂载于 Pod 中的一种存储卷，可以独立于 Pod 资源的生命周期。

7、随书代码， https://github.com/iKubernetes/Kubernetes_Advanced_Practical 。 
