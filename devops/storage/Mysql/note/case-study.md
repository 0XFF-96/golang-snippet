### 为什么 SQL 语句会突然抖动变慢


PS: 暂时看不懂， 14 和 15 章节所讲的内容。   1、当内存数据页跟磁盘数据页内容不一致的时候，这个内存页为 “脏页“。 内存数据写入磁盘后，内存和磁盘上的数据页的内容就一致，称 ”干净页“ 
2、P12 。 
3、  为什么表数据删除一半，表文件大小不变。 
1、innodb_file_per_table , ON。 回收表空间。 
2、数据页的复用跟记录的复用是不同的。 
3、删除数据和插入数据，会造成数据空洞。 （尤其是索引） 4、Online DDL 重建表。 



### Mysql 磁盘碎片化的问题

1、https://qastack.cn/server/202000/how-find-and-fix-fragmented-mysql-tables
2、https://new.qq.com/omn/20201117/20201117A0241F00.html
3、https://cloud.tencent.com/developer/article/1798640
4、https://segmentfault.com/q/1010000023162965
5、https://segmentfault.com/a/1190000038259153
6、[www.cnblogs.com/kerrycode/p/10943122.html](https://www.cnblogs.com/kerrycode/p/10943122.html) 【好文章】  7、free_data_size,


### 如何解决 Mysql 空间不足的问题

学会如何查看空间使用情况。 
可能的原因有几种
1、索引太多导致空间不足。
2、大字段导致空间不足。
3、空闲表空间太多导致空间不足。
4、临时表空间过大，导致空间不足。

ref : https://help.aliyun.com/document_detail/202151.html 。   mysql I/O 高问题
1、https://help.aliyun.com/document_detail/202150.html
2、  mysql 活跃线程数高问题
1、https://help.aliyun.com/document_detail/202148.html 。 
mysql 内存使用问题。
1、https://help.aliyun.com/document_detail/202149.html RDS 定位慢 SQL 的主要方法
1、https://help.aliyun.com/document_detail/202152.html RDS 索引相关问题
1、https://help.aliyun.com/document_detail/52274.html


### 什么是 Using temporary 和 filesort ?

1、Using temporary 和 filesort 这两个是什么意思，有什么优化技巧？

### Mysql 千万大表的优化

1、[www.zhihu.com/question/19719997](https://www.zhihu.com/question/19719997) 2、https://juejin.cn/post/6844903774402641927 3、https://zhuanlan.zhihu.com/p/86137284 

### mysql 如何删除大表

1、mysql optimize huge table, https://stackoverflow.com/questions/5557838/mysql-optimization-of-huge-table 。   2、https://user3141592.medium.com/how-to-scale-mysql-42ebd2841fa6  3、https://avishwakarma.medium.com/mysql-scaling-and-how-to-do-it-d7f5b8dbd5b9 。 水平扩展和垂直扩展。  4、https://avishwakarma.medium.com/mysql-scaling-and-how-to-do-it-d7f5b8dbd5b9
5、[www.oshyn.com/blog/mysql-scaling-options](https://www.oshyn.com/blog/mysql-scaling-options)
6、https://stackoverflow.com/questions/11707879/difference-between-scaling-horizontally-and-vertically-for-databases 7、Horizontal and Vertical Scaling In Databases
[www.geeksforgeeks.org/horizontal-and-vertical-scaling-in-databases/](https://www.geeksforgeeks.org/horizontal-and-vertical-scaling-in-databases/) 。  8、归档策略， https://kebingzao.com/2020/04/01/mysql-archive/ 。 
9、[www.geeksforgeeks.org/horizontal-and-vertical-scaling-in-databases/](https://www.geeksforgeeks.org/horizontal-and-vertical-scaling-in-databases/) 。 


1、https://stackoverflow.com/questions/879327/quickest-way-to-delete-enormous-mysql-table  2、https://stackoverflow.com/questions/879327/quickest-way-to-delete-enormous-mysql-table  3、https://dba.stackexchange.com/questions/53836/how-long-will-it-take-to-rename-table  4、https://blog.mimvp.com/article/30359.html  5、如何删除大表， https://blog.mimvp.com/article/30359.html。   6、https://cloud.tencent.com/developer/news/298548 。  7、https://blog.csdn.net/zyz511919766/article/details/40539333 。  8、http://blog.csdn.net/weixin_33850890/article/details/92345162?spm=1001.2101.3001.6650.2&depth_1-utm_source=distr%E2%80%A6base 。   相关操作：  1、delete
一般用于删除少量表中的数据
优化建议，一定要加上where 条件，并且where条件的列上 一定要有主键或者索引。否则会出现全表扫描的情况
2、drop
直接将表删除，包括表的数据和表的定义。
这种操作，MySQL 会加上一个全局锁，删除期间会阻塞操作，所以对大表来说这样操作一定会对生产造成很大的影响
优化建议，大表不要使用drop table 这种方式，可以考虑先进行 truncate table的方式
3、truncate
这种方式要优于 drop table，风险较之 drop 要小一点，不需要加锁，所以删除表的速度很快
注意事项：这种方式虽然很快，但是需要考虑删除大文件对IO产生的影响  

