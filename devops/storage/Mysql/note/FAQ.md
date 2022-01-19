### Mysql FAQ

1、什么是临时表，临时表什么时候删除？
 2、查询语句不同元素( where join, limit, group by , having 等） 执行先后顺序 
3、有哪些锁（乐观锁悲观锁），select 时怎么加排它锁？
4、mysql的存储引擎innodb中的主级索引和二级索引是指什么？
5、为什么一条 update 语句会锁死整个表？
https://www.cnblogs.com/xiaolincoding/p/15262073.html 
safe update 语句，作为相应的相知



### 关于 MVCC 的理解
1、为什么在 update, (transation 时）， 需要两阶段提交？ 
为什么必须有 “两阶段提交” 呢？这是为了让两份日志之间的逻辑一致。  先写 redo log 后写 binlog， 反之的区别？
 2、磁盘的性能决定了事务提交的性能，也就是数据库的性能， 为什么？ 
3、并发下事务会产生哪些问题、有哪些隔离级别？
 4、可重复读， ABA 问题， 对于当前事务来说直到另一个事务提交之后它再读才会获取到最新结果，
但是它并不知道这期间别的事务对数据做了更新，这就是幻读的问题。  
4、gap-lock 的相关条件，  where 条件必须是非唯一键， where 条件是主键的时候，只会有 record lock 不会有gap lock 。 
5、gap lock 解决的问题， Gap Lock 就是为了防止在本事务还未提交之前，别的事务在当前事务周边插入或修改了数据造成读不一致。
6、Next-key-lock, 相关 Record Lock 和 gap Lock 的组合，  
 
### mysql 相关库

1、如何抽象 msyql 库 （ go-common , daycam ) 
2、一条 sql 是如何被执行的 ？
3、这个库经过哪些重大的 PR 改造？
4、这个库如何被抽象？   代码实践： 1、"http://github.com/jmoiron/sqlx" 利用这个进行实例化。 
2、_ "http://github.com/go-sql-driver/mysql" 3、"database/sql"  4、上面这几个库有什么关系？


### Clustered Index 主键索引和非主键索引的区别

1、聚簇索引和非聚簇索引区别？聚簇索引与非聚簇索引的区别是：叶子节点是否存放一整行记录。如果没有的话，
    可能需要【回表】操作，去聚簇索引树上面，拿记录。
 2、松散索引扫描？https://blog.csdn.net/zm2714/article/details/7887093 
3、go 组件学习， https://mp.weixin.qq.com/s/-2T9BovG8TG32DQKn93LaA 

 