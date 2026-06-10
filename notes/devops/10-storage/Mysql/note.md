### Mysql 的事务机制

1、https://www.cnblogs.com/rickiyang/p/13652664.html 
【五分】 
2、https://www.cnblogs.com/Finley/p/5289161.html 
 3、如何查看锁表的记录和日志？  概念 

1、ACID  2、脏读、不可重复读、幻读
2、四种隔离级别。各自会出现什么效果？
3、三种思路进行并发控制（ 悲观锁、乐观锁、快照）
4、CAS原语是乐观锁的一个典型示例子 
5、mysql 行级锁，GAP 锁 间隙锁、Record lock 单行记录的锁、next-key lock = record lock + next-key lock 。  gap 锁的例子， START TRANSACTION;
 DELETE FROM user WHERE age < 18 （ 锁定了多少行记录？）  2> START TRANSACTION;
2> INSERT INTO user (age) VALUES (17);
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction 。  缺点，  GAP LOCK 锁定粒度相当大，很可能导致吞吐量下降甚至死锁 。 
8、mysql 锁的分类。 1、按照锁定的类型： 表级、行级、页面 2、表，读锁（共享）、写锁（排他）
3、页面， 页面锁 因为MySQL 数据文件存储是按照页去划分的，所以这个锁是 MySQL 特有的。开销和加锁时间界于表锁和行锁之间，锁定粒度界于表锁和行锁之间，并发度一般。    innoDB 如何实现事务。
1、InnoDB引擎中实现事务最重要的东西就是日志文件，保证事务的四大特性主要依靠这两大日志 。   2、redo log 保证 【事务持久性 D】。
3、undo log 回滚日志，保证【事务原子性 A】。
4、一致性 和 隔离性 是通过 MVCC 机制 和 锁机制来一起控制 。   Mysql 的隐事务机制
1、autocommit 参数的修改指只针对当前连接
2、show variables like 'autocommit' 3、autocommit 属性值所影响的特殊指令。 （DDL， create table / drop table / alter table)   undo log 的存储方式。 （隔离行？） 
1、表空间位于共享表空间的回滚段中，共享表空间的默认的名称是 ibdata，位于数据文件目录中 2、理论上当前事务结束，undo log 日志就可以废弃 。 因为由 redo log 处理了。  3、从 undo log 中分析出该行记录以前的数据是什么，从而提供该行版本信息，让用户实现非锁定一致性读取  redo log （tracing-back log) 1、redo log 即重做日志，重做日志记录每次操作的物理修改。用于操作回溯的。    如果没有 redo log, 只有 bin log ?
1、持久性问题一般在发生故障的情况才会重视 。  2、需要假设一个故障层面 （ mysql 层面的故障，而不是程序层面的故障） 3、两种情况， 先更新数据文件，再写入 binlog。 先写入 binlog，再更新数据文件。  4、 crash-safe 的能力。   redo log 和 binlog 的区别
1、binlog 是MySQL server 层实现的功能。 redo log 是 innoDB 引擎特有的 2、redo log 是物理日志， 记录 “在某个数据页做了什么修改”。 binlog 是逻辑日志，记录 sql 语句的原始逻辑 3、redo log 空间是固定的，用完之后会覆盖之前的数据；binlog 是追加写，当前文件写完之后会开启一个新文件继续写 4、redo log 空间是固定的，用完会覆盖之前的数据。 binlog 是追加写， 当前文件写完之后，会开启一个新文件继续写 redo log 两部分。  5、两阶段提交就是为了防止 binlog 和 redo log 不一致发生。

一个更新事务的整体流程。    QA 1、为什么在 update, (transation 时）， 需要两阶段提交？ 为什么必须有 “两阶段提交” 呢？这是为了让两份日志之间的逻辑一致。  先写 redo log 后写 binlog， 反之的区别？ 2、磁盘的性能决定了事务提交的性能，也就是数据库的性能， 为什么？ 3、并发下事务会产生哪些问题、有哪些隔离级别？ 4、可重复读， ABA 问题， 对于当前事务来说直到另一个事务提交之后它再读才会获取到最新结果，但是它并不知道这期间别的事务对数据做了更新，这就是幻读的问题。  4、gap-lock 的相关条件，  where 条件必须是非唯一键， where 条件是主键的时候，只会有 record lock 不会有gap lock 。 5、gap lock 解决的问题， Gap Lock 就是为了防止在本事务还未提交之前，别的事务在当前事务周边插入或修改了数据造成读不一致。
6、Next-key-lock, 相关 Record Lock 和 gap Lock 的组合，   


### REF
1. 什么是幻读？https://www.zhihu.com/question/47007926


