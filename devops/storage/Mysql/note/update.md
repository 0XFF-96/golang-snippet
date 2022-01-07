### 日志系统 update 是如何被执行？
（极客时间的 Mysql) 

1、日志模块, redo logo  （临时日志）、WAL 技术 （ring buffer ) write pos 和 checkpoint 。 
保证 crash-safe ，数据库发送异常重启，提交的记录不会丢。
 redo log 是 innoDB 特有的日志
2、binlog 日志是 server 层， 主要设计是用于归档的。 
3、redo 和 binlog 的三点不同：所属模块不同、记录的内容不同、记录的方式不同。  
【5、利用一条 update 语句，加深对这两个日志的理解。update 语句执行流程图。 】  
6、两阶段提交， prepare 和 commit 两个状态。
是不是意味着会产生两条 redo log 和 一条 binlog ?
  先写 redo log , 后写 binlog  。 和先写 binlog ， 后写 redo log 。 假设在执行事务的每个步骤都会出现奔溃， 
然后想一下此事的恢复方案和状态。   
7、sync_binlog 设置为 1 ， 每次事务的 binlog 都持久到磁盘。   
8、update 响应一次 sql 需要写三次磁盘， redo log 2 两次， binlog 1 次。

### Q & A   

1、如果表 T 中没有字段 k, 而你执行了这个语句 select * from T where k= 1, 产生的报错。是由哪个过程产生的呢？
2、怎么让数据库恢复到半个月内任意一秒的状态。
 3、一条更新语句是如何执行的？ 为什么单到 redo log 或者 binlong 没有办法做到回滚恢复？ 用 binlog恢复到库跟原库相同。 
4、数据库：一周备份和一天一备份 的区别？ 救命的速度和后悔的速度。 恢复时间、恢复点。  
5、如何安全低给小表加字段？关于 MDL 锁 。  ALTER TABLE xxx NOWAIT ADD COLUMN 
6、mysql 如何构造一个 “ 数据无法修改” 的相关场景？
	