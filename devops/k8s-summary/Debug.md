# Debug 问题


### pod debug 

k8s 调度问题。2022-1-18 。 
1、pod 删除不了，也不能被正确调度
2、排查到是 某一个 node 的问题，但是 node 的指标一起正常。cpu， 内存，磁盘。 
3、k8s 的 event 也是一切正常。
 4、kubectl -> api-server 的日志查不到什么有用的信息。 
 5、kubelet, 只能查看每个节点，负责接收来自 k8s api-server 信息的组件。看看这情况是否正常。