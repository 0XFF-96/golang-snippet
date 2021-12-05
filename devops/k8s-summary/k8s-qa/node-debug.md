### k8s-node-debug 

1、内存耗用指标， https://blog.csdn.net/adaptiver/article/details/7084364 。 
2、工作节点数据指标， https://kubernetes.io/zh/docs/tutorials/kubernetes-basics/explore/explore-intro/ 。 
3、集群故障排查， https://kubernetes.io/zh/docs/tasks/debug-application-cluster/debug-cluster/ 。 

q&a
1、为什么会出现很多 evicted 的 pod, 但是有一部分确实正常运行？ 2、evicted pod 会造成什么影响？3、查看节点相关状态， http://blog.csdn.net/wangkaizheng123/article/details/111686685?spm=1001.2101.3001.4242 。 4、节点内存和 cpu 排名， kubectl --context=prod top pods 。 