### jupiter 留下的相关笔记📒

### jupiter 
一、疑惑🤔？
1、服务部署上线健康检查。 
2、对比当前上的服务，和之前发不到一个服务直接日志的
3、服务发布、依赖分析治理。 


api gateway 的相关例子。 
1、https://github.com/zhuCheer/libra

二、监控告警的相关信息。 
1、统一错误码信息
2、错误区分， 区分系统级别的错误和业务级别的错误 
3、区分重要信息和非重要信息。 
4、sop 手册（标准操作手册），这个可以。

三、错误
1、错误收敛和混沌工程。 
2、服务饱和度 和 水位相关的概念。 
方法， 
1、统一错误码信息
2、统一索引信息
3、区分系统和业务。
4、区分重要信息和非重要信息
5、代码即是文档。 
。 错误码 - 》 某一个文件的某一行。 


### 相关？


N 和 M 都是可调参数
1、研究一下 backoff 算法的实现原理。 
2、”http://github.com/cenkalti/backoff"  这个算法的用处在哪里？
3、"http://github.com/go-resty/resty/v2"
( 自己仓库的 http pkg 可以按照这个方向进行优化和改进） 
4、这种技术上实现比较快
疑惑🤔
1、jupiter-agent  的拆分方式，为什么和 jaeger 类似？ 有什么可取之处？
2、jupiter-agent 、juno、Jupiter 三个之间的关系是什么？ （部署、运行、初始化、使用指南） 

### juno 
1、https://github.com/douyu/juno/blob/master/docs/wiki-cn/quickstart.md
2、http://jupiter.douyu.com/juno/ 

二、相关疑惑
1、如何初始化一个基本的 juno ?
2、数据怎么 mock ?
3、数据库的表结构，怎么通过 make install 命令写进去？
有可能是通过 gorm , 以代码转换为 mysql 语句的形式写进去的。 
