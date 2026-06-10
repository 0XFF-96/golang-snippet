### Mysql 锁

1、MDL 作用是？ DDL 和 DML 并发冲突。
2、全局锁， FTWRL 。 全库逻辑备份时需要。
3、表级锁（元数据锁 meta data lock )。进行 DDL 的主要原因。
4、死锁和死锁检测，导致 mysql CPU 100% 。 控制并发度，对于相同行的更新，在进入引擎之前排队。innodb_lock_wait_timeout 
5、InnoDB 在实现 MVCC 时用到的一致性读视图， 用于支持 RC 和 RR， 它没有物理结构，作用是在事务实行期间用来定义 “我能看到什么数据” 
InnoDB 利用了 “所有数据都有多个版本“ 的特性，实现创建快照” 的能力。 

### 锁🔒

1、如何实现 MVCC 机制
2、隔离级别、和提交级别？ 
3、gap 锁的应用场景？ https://blog.51cto.com/u_14916031/3657843
 https://cloud.tencent.com/developer/article/1806998 [www.modb.pro/issue/9277](https://www.modb.pro/issue/9277) https://zhuanlan.zhihu.com/p/73816334 
  4、https://blog.csdn.net/waves___/article/details/105295060 。 什么是 MVCC ？ 
 5、什么是意向锁🔒？ https://tonydong.blog.csdn.net/article/details/103324323


### Mysql 的乐观锁和悲观锁
1、[www.cnblogs.com/kismetv/p/10787228.html](https://www.cnblogs.com/kismetv/p/10787228.html) 
2、https://segmentfault.com/a/1190000016611415 
3、什么是 ABA 问题。 ？ 
4、https://blog.csdn.net/cmm0401/article/details/115816459 。 乐观锁和 MVCC , 实现过 MVCC 。 

