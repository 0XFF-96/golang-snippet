### 简介绍
对于并发操作而言，原子操作是个非常现实的问题。典型的就是i++的问题。 
当两个CPU同时对内存中的i进行读取，然后把加一之后的值放入内存中，可能两次i++的结果，这个i只增加了一次。 
如何保证多CPU对同一块内存的操作是原子的。 golang中sync/atomic就是做这个使用的。

具体的原子操作在不同的操作系统中实现是不同的。
比如在Intel的CPU架构机器上，主要是使用总线锁的方式实现的。 
大致的意思就是当一个CPU需要操作一个内存块的时候，向总线发送一个LOCK信号，
所有CPU收到这个信号后就不对这个内存块进行操作了。 
等待操作的CPU执行完操作后，发送UNLOCK信号，才结束。 
在AMD的CPU架构机器上就是使用MESI一致性协议的方式来保证原子操作。 
所以我们在看atomic源码的时候，我们看到它针对不同的操作系统有不同汇编语言文件

### atomic 的源码实现
1. https://changkun.de/golang/zh-cn/part4lib/ch15sync/atomic/
2. https://docs.kilvn.com/The-Golang-Standard-Library-by-Example/chapter16/16.02.html
3. http://www.poloxue.com/qa/go-automic-load-source-code-question/
4. cpu 缓存协议与 atomic https://studygolang.com/articles/1495

### atomic unit32 做减法的方式
1. http://blog.xyecho.com/timegeekbang-go-sync-atomic/ (极客时间、int32)


### atomic 的练习
1. https://www.geeksforgeeks.org/atomic-adduint32-function-in-golang-with-examples/
2. https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go 
The maximum value for an int type in Go