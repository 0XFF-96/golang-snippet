# 利用 redis 设计消息队列

### 前置知识
- 在看这篇文章之间，你必须要过一遍 BLPOP, RPush Redis 阻塞队列的机制。`BRPOP key [key ...] timeout`
- redis 支持 atomic list operations 
- 

### 
1、http://big-elephants.com/2013-09/building-a-message-queue-using-redis-in-go/
2、开源代码库：https://github.com/adjust/redismq
3、Redis 阻塞队列的机制:https://redis.io/commands/blpop
4、blpop 的相关实现,https://stackoverflow.com/questions/41294931/how-is-brpop-implemented-in-redis 
5、https://redis.io/commands/rpoplpush
6、http://big-elephants.com/2013-10/tuning-redismq-how-to-use-redis-in-go/ 
7、v2 版本的 redis a durable message queue system for go based on redis, see also https://github.com/adjust/rmq
