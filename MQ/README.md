### Golang 相关的各个消息队列研究
1、https://tech.meituan.com/2016/07/01/mq-design.html [美团消息队列设计精要]
解耦、最终一致性、广播、
2、https://juejin.im/post/5cd12fa16fb9a0320b40ec32 [goim 相关文章]
3、基于 go 和 redis 实现轻量消息队列 https://segmentfault.com/a/1190000019692054
4、有赞延迟队列设计思路 [https://tech.youzan.com/queuing_delay/]
5、http://charlesleifer.com/blog/multi-process-task-queue-using-redis-streams/ [py 的可以不看]
6、https://alexstocks.github.io/html/im.html 一套简洁的即时通信(IM)系统
7、https://alexstocks.github.io/html/pubsub.html
8、携程异步消息实践 [https://blog.csdn.net/qiansg123/article/details/80121900]

###  即时通讯技术研究
1、http://www.mit.edu/~6.005/fa09/projects/guichat/assignment.html
2、

### 相关代码库
1、https://github.com/wuzhc/gmq 
2、https://github.com/ouqiang/delay-queue [参考有赞文章设计实现]
3、https://github.com/nats-io/nats-server nats-server 相关
4、https://github.com/alberliu/gim [5 星🌟值得研究]

### Notes

1、我们的消费队列，与其他的消费队列有什么特点？
2、相似的功能之处，不同的功能之处？
3、优缺点？为什么选择这样的一个模式？
faq:
1、衡量队列堆积的指标是什么？
2、是否可以通过开关，从 redis 消费队列，切换到 kafka 或者是 nsq 等其他的消费队列中？
3、会有竞争吗？
4、N 个生产者 -》 一个消费者模式