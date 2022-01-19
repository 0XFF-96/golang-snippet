### 缓存

- 缓存的概念
- 缓存的框架
- 缓存在不同场景下的应用

### 引用
- LRU cache: https://blog.csdn.net/caoshangpa/article/details/78783749

### 相关策略
- LRU 的问题在于，如果在某个数据在前9分钟访问了1万次，
最近1分钟没有访问，那么依然会认为该 key 并不热门而有可能被驱逐。

- LFU 的问题在于，经常会有一些数据在某时刻非常极其热门，
但之后一直没人访问，例如因为某些原因被隐藏的用户动态这类场景。
另外，LFU 的频率信息在缓存失效后依旧会存在内存中。

### 相关情况

1、https://segmentfault.com/a/1190000019222661 【内存泄露】
2、http://juejin.cn/post/6925435360654655495 。 
3、设计高性能本地缓存， https://blog.joway.io/posts/modern-memory-cache/ 。 
4、go 实现缓存组件， https://researchlab.github.io/2016/12/17/go-cache-component/ 。 
5、https://my.oschina.net/slagga/blog/4480099
6、实现缓存系统， https://zhuanlan.zhihu.com/p/81867557 。 
7、go-zero 缓存 https://mp.weixin.qq.com/s/ueo9zNMx1OO4cScebORN9Q 
8、https://coolshell.cn/articles/17416.html
9、Golang 高性能 LocalCache：BigCache 设计与分析
https://pandaychen.github.io/2020/03/03/BIGCACHE-ANALYSIS/

### READ  

#### 如何实现可靠内存缓存？
1、本地内存缓存系统的要求是什么？（如何达到？）
2、简易的 map -> 线程安全的 map -> 高并发的 map -> 无锁的 map -> （如何实现无锁队列？） 
3、LRU 和 LFU 之间有什么区别？在哪些场景下有差异。 访问的频率、访问 key 的分布
4、缓存一定要存在 map[cache_id]cache_value 吗？ 这样做的坏处是什么？有什么弊端没有？

```
文章很不错, 看得出花了不少心思, 多谢分享, 但也有几个问题, 不吝赐教
1. 无论是分段的还是分桶的, 您的文章中value都是设计在map中存放的, 如果是复杂类型, 比如*struct{a,b,c}, key非常庞大的话, 势必要引起海量的gc, 内存也会持续增长.
2. 对于1, 为了解决这个问题, bigcache开了个大buf来保存数据, 使用map[uint64]uint64来存放value存放的位置.
3. 您提到的lfu, lru等等, 以及大小堆, 都是寻找一种淘汰策略, 假设我们找到了一个key A, 在map中存放直接删除就ok了, 但如果我们要解决问题1, 采用了了一个大buf, 这个A有可能在buf的中间位置, 我们删除后中间就是一个内存碎片, 为了合理利用这些碎片, 又要花费一些数据结构来解决

如果考虑以上123, 怎么才能设计出一个比较好的Go本地内存缓存呢

4. 关于环, 为什么不能是直接喂给LFU, 而是一定要走一个ring呢? 这里有不解
```

 https://blog.joway.io/posts/modern-memory-cache/ 【这篇文章值得读两遍】
 知识点： 

### 缓存更新的套路： 

https://coolshell.cn/articles/17416.html 【更新的策略】
 - 写 db or 写 redis ?
 - 有哪几种方式？分别的优缺点是什么，如何构建案例。更新缓存的的
 Design Pattern有四种：Cache aside, Read through, Write through, Write behind caching
 
 问题？（更新和查询 同时） （查询和查询同时）（更新和更新同时） 
 - 缓存的三个操作：更新、查询 （是否命中） 
 两个并发操作，一个是更新操作，另一个是查询操作，更新操作删除缓存后，查询操作没有命中缓存，先把老数据读出来后放到缓存中，然后更新操作更新了数据库
 
 
1、需要考虑的点
- 数据是否一致
- 优缺点，有哪些案例会造成问题
- 复杂度和实现难度


#### go-zero 内存缓存
- safe map 
- 利用 timewheel 实现，可以自定义过期时间的缓存。

### 基于 redis 的多级缓存设计
一、golang基于redis和机器内存的多级缓存设计。 
二、各种常见缓存类型： CDN 缓存、本地缓存等：https://learnku.com/articles/42357

### 缓存优化策略

1、https://tech.meituan.com/2016/12/02/performance-tunning.html  经常需要在代码中控制， gorountine 的生成。
2、goroutine 的并发控制。 worker 模式 。


### to-read 
1. https://dev.to/douglasmakey/how-bigcache-avoids-expensive-gc-cycles-and-speeds-up-concurrent-access-in-go-12bb
2. https://blog.allegro.tech/2016/03/writing-fast-cache-service-in-go.html
3. https://colobu.com/2019/11/18/how-is-the-bigcache-is-fast/ 