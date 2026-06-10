# Golang 小抄 / Cheatsheet

> 每次迭代学习 Go 时，需要重新回顾的知识点合集。
> A collection of Go knowledge points worth revisiting each iteration.

---

## 1. 并发 / Concurrency

### Channel
1. 由【发送者】来关闭 channel，始终如此。
2. 多发送者场景：通过中间代理（middle proxy）通知所有 sender 停止发送，再关闭。
3. channel 的 happen-before 模型。

### Goroutine
1. 常用并发模型：channel 通知控制、`*sync.WaitGroup`、`context.Context`。
2. 协程 vs 线程 vs 进程：从 **内存空间**、**管理的资源**、**调度逻辑** 三个维度对比。
3. 如何限制 goroutine 数量？每次 `go` 之前判断当前数量，超出则阻塞。
4. goroutine 调度时机不确定 → 不能提前 `close` channel。

### Select & 原子操作 / Atomics
1. `select` 机制：`pollorder` 保存 scase 序号，乱序是为了执行时的随机性。
2. CAS（Compare And Swap）：原子操作的一种，著名的无锁算法基础。

### 锁 / Locks
1. Go 中的各种锁，以及它们的源码实现。

---

## 2. 内存与 GC / Memory & GC

1. Go 内存模型中，为什么小对象多了会造成 GC 压力？
2. 逃逸分析（escape analysis）是什么。
3. 内存管理：通常指 **堆内存** 管理，栈内存不需要程序操心。
4. Go 内存分配器：`mspan` 的表示形式，`mcache`，以及 P / mcache / mspan 三者关系。

---

## 3. 运行时与语言机制 / Runtime & Language Internals

### 下划线 `_` 的三种用途
1. `import` 中：触发该 package 的 `init()`。
2. 函数代码中：作为占位符，跳过未使用返回值的编译检查。
3. 全局变量中：编译期检查类型是否实现某接口。
   - 参考：<https://i6448038.github.io/2018/07/18/golang-mistakes/>

### 函数与返回值
1. `return` 不是原子操作，被拆为：**赋值返回值 → 执行 defer → 返回调用方**。
2. 函数返回值通过堆栈返回。

### 抽象层次
- `go user code → go runtime → syscall`，一步步抽象。

### 语言特性
- Go 没有继承、虚方法、重载，只有组合（composition）。

### HTTP 包
- `net/http` 主要由 **`ServeMux`** 与 **`Handler`** 组成。
- `ServeMux` 本质是 HTTP 请求路由器（Multiplexor）：将请求 URL 与预定义路径比对，匹配后调用关联的 Handler。

---

## 4. 常见陷阱 / Common Pitfalls

### For-range
1. 循环永动机 / 神奇的指针 / 遍历清空数组。
2. 不同遍历方式在 [`cmd/compile/internal/gc.walkrange`](https://draveness.me/golang/tree/cmd/compile/internal/gc.walkrange) 中处理。
3. 清空 slice：[`runtime.memclrNoHeapPointers`](https://draveness.me/golang/tree/runtime.memclrNoHeapPointers)。
4. 官方 wiki：<https://github.com/golang/go/wiki/CommonMistakes>。

### 其他常见错误
- 双检查实现单例（double-checked locking）。
- 要求每秒钟调用一次 `proc` 不退出。
- 【Golang 错题本】整体回顾。

---

## 5. 深入主题 / Deep Dives (FAQ)

每次学到新知识点时，可结合曹大的 *golang-notes*、*Introduction-to-Golang* 项目、饶大的源码解析重新回答：

1. 深入理解 Go map：设计与实现。
2. 从源码角度看 Go 的调度机制。
3. Go 程序是怎么跑起来的。
4. Go 的内存管理（分配器细节，见 §2）。
5. Go 的垃圾回收机制。
6. Goroutine 协程实现。
7. Go heap 源码。
8. `sync` 包的设计与实现。

---

## 6. 资源 / Resources

### 文章 / Articles
- [Golang 的 GC 和内存逃逸](https://blog.csdn.net/qq_35976351/article/details/104602154)
- [data-structures-questions — Golang 部分](http://github.com/KeKe-Li/data-structures-questions/blob/master/src/chapter05/golang.01.md)
- [饶全成 — interface 解析](https://www.cnblogs.com/qcrao-2018/p/10766091.html)
- [draveness — 反射](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/)
- [draveness — 2020 总结](https://draveness.me/2020-summary/)

### 仓库 / Repos
- [tmrts/go-patterns](https://github.com/tmrts/go-patterns)
- [golang/go wiki — Books](http://github.com/golang/go/wiki/Books)
- [golang/go wiki — CommonMistakes](https://github.com/golang/go/wiki/CommonMistakes)

### 提交 / Commits
- [select 相关 commit 5038792](https://github.com/golang/go/commit/5038792837355abde32f2e9549ef132fc5ffbd16)
