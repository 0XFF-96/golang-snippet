### map 遇到过的坑
1. 并发读写问题


### 想看的有关 map 的文章
1、https://blog.golang.org/maps (基础官方文章肯定要过一遍)
2、https://www.jianshu.com/p/f2e7650da938
3、delete key in map: https://stackoverflow.com/questions/1736014/delete-mapkey-in-go

### 实现 map 
1. 如何设计并实现一个线程安全的 Map ？(下篇)
https://halfrost.com/go_map_chapter_two/

2. 如何设计并实现一个线程安全的 Map ? (上篇)

- 选择一个优秀的哈希算法，用链表 + 数组 作为底层数据结构，如何扩容和优化，这些应该都有了解了。

### map 的相关 issue 
1. https://github.com/golang/go/issues/20135 []
runtime: maps do not shrink after elements removal (delete) 。内存泄露
这里可以参考 go-zero 的解决方案。 


### 为什么 go key 是无序的？
- https://www.bookstack.cn/read/qcrao-Go-Questions/map-map%20%E4%B8%AD%E7%9A%84%20key%20%E4%B8%BA%E4%BB%80%E4%B9%88%E6%98%AF%E6%97%A0%E5%BA%8F%E7%9A%84.md


### map 相关阅读

1、https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/
2、https://segmentfault.com/a/1190000020616487
3、https://juejin.cn/post/6844903848587296781#heading-0

### 问题
1. fatal concurrent map read and write. https://github.com/golang/go/issues/26010 
2.  