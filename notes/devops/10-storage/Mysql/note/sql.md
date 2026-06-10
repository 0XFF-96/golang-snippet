# SQL 上的相关优化

### 分段统计


case-when , 相关。 
1、http://segmentfault.com/q/1010000009539927 ，  2、https://blog.csdn.net/qq_33562122/article/details/52939464 3、https://zhuanlan.zhihu.com/p/409669710 4、怎么可以优化 case when 写太多的问题？ 5、https://blog.csdn.net/qq_33562122/article/details/52939464。 6、https://developer.aliyun.com/article/395746 。 最通用的统计。  7、[www.w3schools.com/mysql/func_mysql_case.asp](https://www.w3schools.com/mysql/func_mysql_case.asp)  。  8、如何进行行列转换？  有一张考试成绩表，主要有学号stu_id，成绩s_score字段，学生成绩在[0, 100]的区间中，现在按5分为一个分数段，统计各个分段中学生数量的分布情况。   Case Example: 1、SELECT CustomerName, City, Country
FROM Customers
ORDER BY
(CASE
    WHEN City IS NULL THEN Country
    ELSE City
END);  2、Example2   CASE
    WHEN condition1 THEN result1
    WHEN condition2 THEN result2
    WHEN conditionN THEN resultN
    ELSE result
END;  


### 如何按照一周的区间 group by 

1 、 https://dba.stackexchange.com/questions/164598/group-by-week-but-show-a-more-readable-the-result  2、SELECT DATE_SUB(date, INTERVAL WEEKDAY(date) DAY) AS wk 
FROM user_passive_behavior_daily_stat WHERE user_id = 335 AND date >= "2020-09-18 11:03:06" GROUP BY wk ORDER BY wk desc
 
This section describes the functions that can be used to manipulate temporal values. See [Section 11.2, “Date and Time Data Types”](https://dev.mysql.com/doc/refman/8.0/en/date-and-time-types.html), for a description of the range of values each date and time type has and the valid formats in which values may be specified.

### 增量同步

1、https://blog.csdn.net/weixin_42184548/article/details/113208238 【mysql 主从同步机制】  2、

canal 原理
1、模拟 mysql slave 的交互协议，伪装成 mysql slave, 向 mysql master 发送 dump 
2、https://segmentfault.com/a/1190000038920200   实时数据同步方案。  1、Mysql 中的数据同步到 Hive 来进行数据分析 。  2、[www.infoq.cn/article/g7r4bmk3nadpgygyyild](https://www.infoq.cn/article/g7r4bmk3nadpgygyyild) 3、


3、https://github.com/go-mysql-org/go-mysql 。  

2、[www.jianshu.com/p/5e6b33d8945f](https://www.jianshu.com/p/5e6b33d8945f) 3、深入 mysql replication 协议。  4、https://github.com/go-mysql-org/go-mysql-elasticsearch 。 自动同步。 



