### 使用 stackdriver 生成日志

1、https://kubernetes.io/zh/docs/tasks/debug-application-cluster/logging-stackdriver/
2、阅读此篇前，必须要读 ： https://kubernetes.io/zh/docs/concepts/cluster-administration/logging/ 。 

2、kubectl describe node $NODE_NAME
3、kubectl label node minikube http://beta.kubernetes.io/fluentd-ds-ready=true.  【打上 label 】
Note : 如果节点发生故障并且必须重新创建，则必须将标签重新打在重新创建了的节点。 为了让此操作更便捷，你可以在节点启动脚本中使用 Kubelet 的命令行参数给节点添加标签。 
4、kubectl apply -f https://k8s.io/examples/debug/fluentd-gcp-configmap.yaml
部署一个带有日志代理配置的 configmap 
5、kl get ds --all-namespaces 。 检查是否有运行成功。 
6、kubectl apply -f https://k8s.io/examples/debug/counter-pod.yaml
【运行一个无限打 log 的 counter 服务】
7、kubectl get pods 、logs -f 查看是否打 log 成功。 

resource: 
1、https://cloud.google.com/bigquery/ 【big query 服务】

