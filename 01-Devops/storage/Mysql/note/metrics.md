### Mysql 性能关键指标

1、mysql.cpu_usage,  服务进程 CPU 使用率 
2、mysql.tb.tmp.disk， 临时表数量
3、mysql.inodb_data_fsyncs ， innoDB 平均每秒 fsync 操作次数。 
4、mysql.iops, Mysql 读写次数。 

### Merge-worker 

1、为什么指标出现异常动态
2、mysql 的 cpu 上升取决于什么？
3、有哪些指标异常？

具体看看这个 PR, 
1、  https://github.com/sundayfun/daycam-server/pull/5686  。 
2、  https://github.com/sundayfun/daycam-server/pull/5671   。 

