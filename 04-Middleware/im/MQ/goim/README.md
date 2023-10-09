### 知识点

1、Round: goim 自己进行了buffer的管理，避免了频繁申请内存的开销。
通过自定义的Pool结构来分配Buffer，因为分配时要加锁，使用Round来组合多个Pool，
通过mod运算随机获取一个Pool，来减缓锁的争用
Pool内部使用一条单链表，维护一个free指针指向未分配的 buffer 

2、Ring是Channel内部用来保存并重用 proto 的一个结构,  
  
3、Bucket是Session的容器，为了减少锁争夺，会有多个Bucket，根据用户id与Bucket数量进行mod运算来确定，

这个Session放到哪个Bucket中，是一种很常见的sharding

4、Router 是旧版的代码才有的特性，详细看这里

> https://github.com/TitanBenny/goim/tree/master/router 


### Reference 
1、https://moonshining.github.io/tags/goim%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/
2、https://juejin.im/post/5cd12fa16fb9a0320b40ec32