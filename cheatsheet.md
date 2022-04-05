# Golang 小抄

### Golang 错题

1、下划线在import中， 目的是为了导入那个 package 的 init.go 函数
2、在函数代码中，目的是为了，占位符，跳过编译检查。 例如，一个函数有返回 3 个参数，但是只用到两个，第三个就用下划线作为占位符。 
3、全局变量中，为了检查这个类型有没有实现相关接口的函数，在编译阶段。 1、https://i6448038.github.io/2018/07/18/golang-mistakes/ 。 
4、【golang 错题本】

### 关于 channel 

1、由写的人来关闭 channel ?
2、始终由 【发送者】来关闭 channel 
3、中间代理，通知所有 xxxx sender 

### golang 相关

1、golang 中常用的并发模型。 chnnel 通知控制、waitGroup（*sync.WaitGroup）、ctx 实现并发控制。
2、协程、线程、进程的区别。 从三个方面来回答吧， 内存空间、管理的资源、调度的逻辑。 
3、Golang的内存模型中为什么小对象多了会造成GC压力？
4、go user code -> go runtime -> syscall ， 一步步抽象成相关的代码。
5、golang 中的锁，以及他们怎么被实现？源码解读？
6、如何限制 goroutine 的数量？所以我们可以限制下Goroutine的数量,这样就需要在每一次执行go之前判断goroutine的数量，如果数量超了，就要阻塞go的执行。
7、Go 的函数返回值是通过堆栈返回的, return 语句不是原子操作，而是被拆成了两步。 函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
8、golang 中的 select 机制， pollorder保存的是scase的序号，乱序是为了之后执行时的随机性。
9、CAS算法（Compare And Swap）,是原子操作的一种, CAS算法是一种有名的无锁算法
10、Go中的逃逸分析是什么。
11、golang 中的内存管理。 通常在Golang中,当我们谈论内存管理的时候，主要是指堆内存的管理，因为栈的内存管理不需要程序去操心。
12、Go中的http包的实现原理？ServeMux 本质上是一个 HTTP 请求路由器（或者叫多路复用器，Multiplexor）。它把收到的请求与一组预先定义的 URL 路径列表做对比，然后在匹配到路径的时候调用关联的处理器（Handler）。
Golang中http包中处理 HTTP 请求主要跟两个东西相关：ServeMux 和 Handler。

13、GMP 的相关缺点？其他
1、goroutnine 的调度时机不缺点，不能提前 close channel 
2、for-range 语句中，循环遍历的服用问题。
3、golang 中没有继承，只有组合，也没有虚方法，没有重载如。4、要求每秒钟调用一次 proc 不退出。5、双检查实现单例？。
6、channel 的 happen before 模型。



### Golang 资源

这里面收集的一些问题是，每次迭代学习 go 的时候，都需要重新学习和重新回答一遍的。
因为每次学到新的知识点的时候，曹大的 golang-notes、
Introduction-to-Golang 项目、饶大的 go 相关源码、
 
[FAQ]

1、深入理解 go map ，从设计和实现上。
2、从源码角度，看 golang 的调度机制
3、go 程序是怎么跑起来
4、go 语言的内存管理。Go内存分配器中mspan的表示形式、mcache、Go中P，mcache和mspan之间的关系。 
5、go 语言的垃圾回收机制
6、go 语言之 gorountine 协程
7、go heap 源码
8、go sync 的设计与实现

【资源】

1、https://blog.csdn.net/qq_35976351/article/details/104602154 。 Golang的GC和内存逃逸
2、quick-note, http://github.com/KeKe-Li/data-structures-questions/blob/master/src/chapter05/golang.01.md 。 
3、绕全成，关于 interface 的解析。 [www.cnblogs.com/qcrao-2018/p/10766091.html](https://www.cnblogs.com/qcrao-2018/p/10766091.html) 【interface 】
4、反射， https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/ 
5、https://github.com/tmrts/go-patterns


### For-range 

1、for-ranage 。 循环永动机、神奇的指针、遍历清空数组、[runtime.memclrNoHeapPointers](https://draveness.me/golang/tree/runtime.memclrNoHeapPointers) 
2、不同的遍历方式， 它们会在 [cmd/compile/internal/gc.walkrange](https://draveness.me/golang/tree/cmd/compile/internal/gc.walkrange)  
3、common-mistakes , for-range 。 https://github.com/golang/go/wiki/CommonMistakes 
1、golang 相关的 book , http://github.com/golang/go/wiki/Books 
2、select 相关， https://github.com/golang/go/commit/5038792837355abde32f2e9549ef132fc5ffbd16  
3、https://draveness.me/2020-summary/ 【path】

