### sync.pool 的相关文章

### 等待确定文章来源👇


### sync.Pool作用

对象重用机制,为了减少GC,sync.Pool是可伸缩的，并发安全的

两个结构体

```go
type Pool struct {
    local     unsafe.Pointer // local fixed-size per-P pool, actual type is [P]poolLocal
    localSize uintptr        // size of the local array
    // New optionally specifies a function to generate
    // a value when Get would otherwise return nil.
    // It may not be changed concurrently with calls to Get.
    New func() interface{}
}
// Local per-P Pool appendix.
type poolLocal struct {
    private interface{}   // Can be used only by the respective P.
    shared  []interface{} // Can be used by any P.
    Mutex                 // Protects shared.
    pad     [128]byte     // Prevents false sharing.
}
```
Pool是提供外部使用的对象,Pool有两个重要的成员,local是一个poolLocal数组,localSize是工作线程的数量( runtime.GOMAXPROCS(0)),Pool为每个线程分配一个poolLocal对象

# 写入和读取

* Pool.Get
先获取当前线程私有值(poolLocal.private)获取<br>
否则则从共享列表(poolLocal.shared)获取<br>
否则则从其他线程的共享列表获取<br>
否则直接通过New()分配一个返回值<br>
* Pool.Put
当前线程私有制为空,赋值给私有值<br>
否则追加到共享列表<br>

# sync.Pool注意点
临时性,当发生GC时,Pool的对象会被清除,并且不会有通知<br>
无状态,当前线程中的PoolLocal.shared的对象可能会被其他线程偷走<br>

# 大规模Goroutine的瓶颈
会对垃圾回收(gc)造成负担,需要频繁的释放内存<br>
虽然goroutine只分配2KB,但是大量gorotine会消耗完内存,并且gc也是goroutine调用的<br>

# 原理和作用
原理类似是IO多路复用,就是尽可能复用,池化的核心优势就在于对goroutine的复用。
此举首先极大减轻了runtime调度goroutine的压力，其次，便是降低了对内存的消耗

![image](https://user-images.githubusercontent.com/87457873/128817600-2399ea3a-a476-4172-9fa3-d973fec904cb.png)

