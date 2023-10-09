### 消息队列学习
1. 当我们谈到消息队列时，一般需要有什么诉求呢？ 如果我们不想用 kafka 这么重的中间件，
有什么办法能为我们业务服务，一些简单定量的中间件？用于解耦？

2. delay-queue 的状态流转图。

3. 消息队列的设计
- producer 
- consumer 
- storage 
- 如何进行数据的存储？ pack、unpack 

- 是否有重试功能？ 如何进行重试？
- 在那个步骤进行 ACK 操作？
- 需要哪些监控 metric ? 如何知道队列是否有堆积、队列堆积延迟多少、队列消费速度？
- How do you scale throughput?
- 

