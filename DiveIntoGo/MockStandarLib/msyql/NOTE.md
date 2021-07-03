### 官方文档笔记
1. an *sql.DB, it should be possible to share that instance between
     multiple goroutines, without any extra synchronization
2.
### 源码学习笔记
1. Driver interface, Driver is the interface that must be implemented by a database

2. Conn interface, Conn is assumed to be stateful.
        And not used concurrently by multiple goroutines.
3. 什么时候调用 Conn.Close() 方法？

```
 connections and only calls Close when there's a surplus of idle connections
```
4. 值得学习的测试技巧，hi jack time.Now() 函数
5. 代码函数调用关系图
```
// nowFunc returns the current time; it's overridden in tests.
var nowFunc = time.Now
```