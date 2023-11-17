
一、Developing and debuging 1、kubectl is configured to communicate with the cluster
2、[Telepresence](https://www.telepresence.io/reference/install) is installed ， http://www.telepresence.io 
3、FAST, LOCAL DEVELOPMENT FOR KUBERNETES AND OPENSHIFT MICROSERVICES， 是一个在本地能够帮助远程 debug 的工具🔧。
二、Application Introspection and Debugging 
1、kubectl apply -f https://k8s.io/examples/application/nginx-with-request.yaml
2、kubectl describe pod nginx-deployment-1006230814-6winp
3、kubectl scale --current-replicas=2 --replicas=1 deployment/nginx-deployment 【 scale 】
4、kubectl get events
5、kubectl get events --namespace=my-namespace

三、debugging a down/unreachable node 
1、kubectl get nodes 
2、kubectl get node -o yaml
3、kubectl describe node kubernetes-node-861hwording: 
1、The container state is one of Waiting, Running, or Terminated 。 Ready tells you 。 Restart Count tells 。 Lastly， 
2、A common scenario that you can detect using events is xxx 
3、 if we requested 1000 millicores then none of the Pods would be able to schedule 。