# 还没有看的一些文章

### 索引分析
1、https://tech.meituan.com/2014/06/30/mysql-index.html
2、[www.cnblogs.com/zsql/p/13808417.html](https://www.cnblogs.com/zsql/p/13808417.html)
3、[www.infoq.cn/article/ojkwyykjoyc2ygb0sj2c](https://www.infoq.cn/article/ojkwyykjoyc2ygb0sj2c)
4、[www.huaweicloud.com/articles/9954d52909f3329014272fddd2784c31.html](https://www.huaweicloud.com/articles/9954d52909f3329014272fddd2784c31.html)
5、https://blog.csdn.net/mysteryhaohao/article/details/51719871


### 一条 SQL 的执行过程详解

1、[www.cnblogs.com/mengxinJ/p/14045520.html](https://www.cnblogs.com/mengxinJ/p/14045520.html)
2、https://blog.csdn.net/weter_drop/article/details/93386581

### Delete 语句引起的相关问题

1、https://qastack.cn/server/202000/how-find-and-fix-fragmented-mysql-tables
2、https://new.qq.com/omn/20201117/20201117A0241F00.html
3、https://cloud.tencent.com/developer/article/1798640
4、https://segmentfault.com/q/1010000023162965
5、https://segmentfault.com/a/1190000038259153
6、[www.cnblogs.com/kerrycode/p/10943122.html](https://www.cnblogs.com/kerrycode/p/10943122.html) 

7、表空间索引， https://help.aliyun.com/knowledge_detail/57848.html?spm=a2c52.12493884.help.dexternal.238797daNk0Y1o 。 
8、https://dev.mysql.com/doc/refman/8.0/en/optimize-table.html
9、 fragment 。 https://lefred.be/content/overview-of-fragmented-mysql-innodb-tables/ 。 

### mysql 死锁日志

1、学会看 mysql 死锁日志。 
2、https://github.com/sundayfun/daycam-server/pull/5816 。   Ref:  1、https://segmentfault.com/q/1010000004272136 。  2、https://database.51cto.com/art/201910/604421.htm 3、https://lotabout.me/2020/God-Damn-MySQL-Locks/ 4、https://learnku.com/articles/39212?order_by=vote_count& 。  5、https://zhuanlan.zhihu.com/p/94778920 。 

### Mysql 死锁相关资料

view, 刚刚本地开1000个Goroutine作并发测试，
发现gap锁和插入意向锁两者一起很容易造成死锁。
代码中事务开启之后只读取一遍数据，
无需考虑重复读取数据不一致的问题。
所以在代码层面降低了事务隔离级别，从可重复读降低到了读已提交后，
并发测试通过 。 
https://github.com/sundayfun/daycam-server/pull/5824  。 
 行锁，意向锁， gap 锁。 
  https://github.blog/2021-09-27-partitioning-githubs-relational-databases-scale/ 
对关系数据库进行分区来处理伸缩 GitHub介绍其如何基于MySQL处理数据伸缩能力的文章。
GitHub没有对表做横向拆分，而是依据业务逻辑纵向拆分。
比较有意思的是通过ProxySQl实现写入切换的过程。
过程中会有短时停机，但能保证不丢失业务数据。  

1、https://blog.csdn.net/u012611878/article/details/81437622 

### For update 语句的使用

When using MySQL's FOR UPDATE locking, what is exactly locked?

1、https://stackoverflow.com/questions/6066205/when-using-mysqls-for-update-locking-what-is-exactly-locked  

2、https://dev.mysql.com/doc/refman/8.0/en/innodb-locking-reads.html

### Mysql 并集和交集

1、https://blog.csdn.net/sanzhongguren/article/details/76615464
2、union 和 interset ?
3、https://blog.csdn.net/sanzhongguren/article/details/76615464 。

### explain 的原理 
 
1、https://dev.mysql.com/doc/refman/8.0/en/explain-output.html  2、





