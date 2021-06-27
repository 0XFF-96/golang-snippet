### 什么是 config-map 


### concept 
ConfigMap和Secret将相应的配置信息保存于资源对象中，而后在Pod对象上以存储卷的形式将其挂载并加载相关的配置，降低了配置与镜像文件的耦合关系。
 
ConfigMap用于在运行时将配置文件、命令行参数、环境变量、端口号以及其他配置工件绑定至Pod的容器和系统组件。Kubernetes借助于ConfigMap对象实现了将配置文件从容器镜像中解耦，从而增强了工作负载的可移植性，使其配置更易于更改和管理，并防止将配置数据硬编码到Pod配置清单中


### Op 
1、k8s 内的 pod 是如何连接 mysql 的？（过程）

2、如何获取 config-map ? 

```kubectl get configmaps demoapp-config -o yaml ```

3、如何通过环境变量引用ConfigMap键值？

4、如何通过 config-map 和 volume 进行文件挂载？

``` kubectl exec configmaps-volume-demo -- ls /etc/nginx/conf.d ```
