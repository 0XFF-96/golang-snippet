### 知识索引

### Chapter1: 锁
    - 程序发生死锁
    - 程序发生活锁
    - 程序发生饥饿
    - 产生上面的情况的条件？
    - 针对具体场景如何解决？

### Chapter3: 
    - context-switch 的性能消耗？如何 bench mark
    - 什么是并发编程的 fork-join model ？
    - 如何知道是什么 goroutine 消耗的内存？
    - Gorountine 内变量引用空间的问题，画出程序执行流程图
    - Mutex, RWMutex 的性能开销
    - cond (不懂...相关条件的使用）
    - once 的相关用法
    - wait group 用法
    - channel 
    - 使用 sync.Pool 减少对象的分配
    
### Chapter4 
    - confinement 限制
    - preventLeak , 如何限制泄露
    - channel 的 error handler 
    - pipeline 怎么区分流处理和批处理 （ discrete or stream / is preemptable?)
    - generator 的写法
    - repeat 函数的相关写法 
    - Little's Law， 如何衡量一个 queue 的质量

### 结论 and Prompts 
 - 	wg.Add(1) 一定要发送在 goroutine 之外，避免调用时机的不确定引起

### Must-to-Read 
- https://en.wikipedia.org/wiki/Little%27s_law  [设计消息队列时，需要考虑的问题]
- https://www.youtube.com/watch?v=hqaSbAykV_Y
- https://www.process.st/littles-law/

### 有趣的文章
http://www.singchia.com/2018/01/29/Concurrency-Patterns-Summary-And-Implementation/

这篇文章，以 Accept tcp 连接为例子，从四个方面：
1、unlimited concurrency
2、limited concurrency
3、elastic concurrency
4、precise concurrency

介绍了 go concurrency 的并发编程模式。通过不断地
迭代和改进示例代码，大大地减少了理解负担，是一个很不错的
学习资料。如果在这个基础上，能够再增加一些关于 channel 的使用就更好了