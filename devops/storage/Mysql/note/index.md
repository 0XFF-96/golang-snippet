### Mysql 的索引


1、mysql 的索引模型、索引数据模型、索引的维护和开销。
2、什么情况下需要 force index 
3、什么情况会出现 【回表】？ 覆盖索引如何解决？
4、什么是【索引下推】？index condition pushdown 。目的是为了减少回表的次数。 


### ？ 
1、这两种不同的索引对于
2、change buffer 是什么？唯一索引不能用。 3、唯一索引的性能消耗， 因为不能用 change buffer ， 会导致内存命中率降低，当存在大量插入操作时。 
4、普通索引/唯一索引的性能在查询上没有什么区别，但是在更新上有区别。 
5、redo log  和 change buffer 之间的差别？

### 为什么 Mysql 会选错索引？

1、force 索引有什么作用？
2、delete from t , 有可能会导致 【索引失效】。
3、shot index 可以看到索引基数。
4、通过 analyze table 命令，用来冲洗统计索引信息。 
6、通过 seelct count(distinct email)  as L from SUser ， 依次选取不同长度的前缀来看这个值。
7、select count(distinct left(email, 4)) as L4



