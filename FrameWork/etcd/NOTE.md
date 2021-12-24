### NOTE - tinyKv 

### ETCD playgroud 

1、https://github.com/etcd-io/etcd-play  
 2、https://github.com/etcd-io/etcd  

### 专注

1、https://hustport.com/d/238-talent-plan-20-tinykv 
 2、描述， https://hustport.com/d/236-talent-plan-20-tinykv 
3、https://hustport.com/d/234-talent-plan-20 
4、开发素养， https://hustport.com/d/228 5、南京大学，开发配置。 
https://nju-projectn.github.io/ics-pa-gitbook/ics2019/PA0.html 


### 资料

1、华中科技大学的讨论区。  1推荐我们学校搞的一个专区 https://hustport.com/t/talent-plan 
②每周五会进行所有战队学习进度的统计并在群内公示
③学习营组织方面的建议和问题，可以向徐珍和宋可欣提出
④除了“方法如何实现？某某数据结构如何定义？”等代码级的问题，其余问题大家可以在群内积极交流讨论，也可以在asktug（https://asktug.com）上 “学习与认证” 板块， Talent Plan 标签下提出、解答互助。
（注意标题下面选择“学习认证”、“Talent Plan”标签）
⑤第一次分享课时间为11月17号晚上8:00到9:40，zoom会议连接：https://pingcap.zoom.us/j/99814310864
zoom会议号：998 1431 0864
期待大家准时参加哟~  3、做了cs186 的b+树，实现还是比较简单 。
  4、Talent Plan 2021 kv 学校营地信息总览。   
5、提问题的方式，   
6、如何提交相关问题？  
7、如何？https://learn.pingcap.com/learner/talent-plan/implement-a-mini-distributed-key-value-database  8、如何解析， https://asktug.com/t/topic/242996/2 。   
9、课程， https://learn.pingcap.com/learner/course/390002 。 
 10、https://en.pingcap.com/blog/tidb-internal-data-storage 。 tikv 技术博客。   
11、数据库 cool 谈， [www.infoq.cn/video/wg7VNmeKXjS4PzMgDwRS](https://www.infoq.cn/video/wg7VNmeKXjS4PzMgDwRS) 。 
  12、raft 典型场景分析， https://zhuanlan.zhihu.com/p/27970722 。 
  13、raft 作者的视频， [YouTube video (vYp4LYbnnW8)](https://www.youtube.com/watch?v=vYp4LYbnnW8&t=986s&ab_channel=DiegoOngaro) 。 
  14、Raft 在 etcd 中的实现
， https://blog.betacat.io/post/raft-implementation-in-etcd/ 。 Go 夜读主推。 
  15、数据库管理系统， https://oceanbase-partner.github.io/lectures-on-dbms-implementation/ 。  
16、https://open.oceanbase.com/competition/index 16、学一门系统编程语言。

### 如何提交作业？

1、不错的讲解，加上视频。  https://hardcore.feishu.cn/docs/doccnMRVFcMWn1zsEYBrbsDf8De# 
raft 算法论文导读， 与 etcd 源码解读。   

2、[www.bilibili.com/video/BV1CK4y127Lj?p=2](https://www.bilibili.com/video/BV1CK4y127Lj?p=2)  这篇文章中有一个和当前教程不一样的结论
[www.jianshu.com/p/3c6a4fd6a7cc](https://www.jianshu.com/p/3c6a4fd6a7cc)
raft协议中有一个约定，不能提交之前任期内log entry作为commit点。这是为什么？
这个问题主要是raft协议中commiting entries from previous term部分看的时候有点困惑，
开始误解成了这个约定是用来保证之前任期内已经被复制到大多数server却没有被提交的日志在新的仍期内不会被覆盖。
实际上，论文中的figrure8的过程是一个正确的过程。
在（c）中,index=2并没有被提交，在(d)中被复制了是一个正确的做法。
论文想阐述的是：如果在(c)中，leader提交了这个之前任期内的entry，
在(d)中依然会被覆盖，也就是说被commit的entry覆盖了，
这是一个错误！因此约定“can not commit entries from previous term”   
3、etcd raft 库解读。 





 






