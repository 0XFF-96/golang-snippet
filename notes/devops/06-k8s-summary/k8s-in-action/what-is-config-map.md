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



### 练习 Config-map 

> http://mp.weixin.qq.com/s#rd?__biz=MzUzNTY5MzU2MA==&mid=2247486416&idx=1&sn=20d568f93d0f39e0f3c7ef3ce42ac1d8&chksm=fa80d%E2%80%A60b05&scene=178&cur_album_id=1394839706508148737 

操作： 
```kubectl create configmap demo-configmap --from-file=configmap-demo ```。 【这种方式不成功】

- 以下这种方式才成功。 kubectl create configmap configmap-demo-2 \
–	--from-file=configmap-demo-game \
–	--from-file=configmap-demo-ui

2、直接用命令行创建 configmap,  ```kubectl create ConfigMap special-config –from-literal=special.level=very –from-literal=special.type=charm ``` 【不行】

3、用 ConfigMap 中的数据定义容器环境变量
–	修改 configmap-pod.yaml 文件里面的相关参数
–	kubectl create -f configmap-pod.yaml 
–	kubectl logs -f xxx , 看是否能够查看容器 env 的相关输出。 
–	sh -c env 

4、将 ConfigMap 挂载到数据卷。
–	在容器中，可以通过目录 /etc/config , [ "/bin/sh", "-c", "ls /etc/config/" ]
–	kubectl create -f pod-configmap-volume.yaml
–	kubectl  logs -f dapi-test-pod-2 , 产生输出如下， configmap-demo-game、configmap-demo-ui
–	ConfigMap引用配置中 使用path 字段为特定的 ConfigMap 项目指定预期的文件名 。 