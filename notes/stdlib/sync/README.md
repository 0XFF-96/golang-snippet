### 文章推荐
1、https://www.jianshu.com/p/ce1553cc5b4f 【sync 进化方向】
2、sync map 原来剖析 
https://jiajunhuang.com/articles/2021_03_15-go_sync_map.md.html

### Sync bug 
1. https://github.com/golang/go/issues/41011
2. https://github.com/golang/go/pull/41000
3. https://gocn.vip/topics/10860

### Sync map 
from:https://cloud.tencent.com/developer/article/1539049
参考：
1. https://juejin.im/post/6844904100287496206 
2. 针对 sync map 的 curd 详细流程图：https://developer.aliyun.com/article/741441

官方解析在以下两种情况下使用 sync.map 会更好。 

```
// The Map type is optimized for two common use cases: 
(1) when the entry for a given key is only ever written once but read many times, as in caches that only grow,

(2) when multiple goroutines read, write, and overwrite entries for disjoint
sets of keys. In these two cases, use of a Map may significantly reduce lock
contention compared to a Go map paired with a separate Mutex or RWMutex.
```

### usage 
1. stabel key 
2. disjoint stores 
3. concurrent loops 

1. 并发安全的 map 
2. 用于针对 value 是 `interface{}` 类型的场景下
3. 针对读多写少多场景
4. (没有实现 java ) 分段🔒相关多技术，


⚠️注意：A Map must not be copied after first use.

### 优点
sync.Map的实现有几个优化点，这里先列出来，我们后面慢慢分析。

1. 空间换时间。通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。
2. 使用只读数据(read)，避免读写冲突。
3. 动态调整，miss次数多了之后，将dirty数据提升为read。
4. double-checking。
5. 延迟删除。删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据。
6. 优先从read读取、更新、删除，因为对read的读取不需要锁。


### 实现的知识点:
1. 利用空间换时间的操作，dirty map 和  read map 
2. 双检查，double-check
3. 

### 疑惑🤔 
1. read.amended 这个数据是怎么被修改的呢？因为m.dirty也指向这个entry,所以m.dirty也保持最新的entry
2. 没有 Len 方法， 
sync.Map没有Len方法，并且目前没有迹象要加上 (issue#20680),所以如果想得到当前Map中有效的entries的数量，需要使用Range方法遍历一次， 比较X疼


### java 对于共享 map 的解决方案
1、jav a的 ConcurrentHashMap的实现，在map的数据非常大的情况下，
一把锁会导致大并发的客户端共争一把锁，Java的解决方案是shard, 内部使用多个锁，
每个区间共享一把锁，这样减少了数据共享一把锁带来的性能影响，

### pool 相关知识点
1. http://cbsheng.github.io/posts/golang%E6%A0%87%E5%87%86%E5%BA%93sync.pool%E5%8E%9F%E7%90%86%E5%8F%8A%E6%BA%90%E7%A0%81%E7%AE%80%E6%9E%90/

### mutex 相关
1. https://colobu.com/2018/12/18/dive-into-sync-mutex/