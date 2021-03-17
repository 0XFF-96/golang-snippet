### golang 面试问题？
转自知乎：记录下我马上直接想到的问题，算给自己出面试题做个笔记
如果要求是CS背景良好、写过20-30万行商用后台代码、1-2年go经验的，我问如下问题：
1. 1.9/1.10中，time.Now()返回的是什么时间？这样做的决定因素是什么?
 2. golang的sync.atomic和C++11的atomic最显著的在golang doc里提到的差别在哪里，如何解决或者说规避？
 3. 1.11为止，sync.RWMutex最主要的性能问题最容易在什么常见场景下暴露。有哪些解决或者规避方法？
 4. 如何做一个逻辑正确但golang调度器(1.10)无法正确应对，进而导致无法产生预期结果的程序。调度器如何改进可以解决此问题？
 5. 列出下面操作延迟数量级(1ms, 10us或100ns等)，cgo调用c代码，c调用go代码，channel在单线程单case的select中被选中，high contention下对未满的buffered channel的写延迟。
 6. 如何设计实现一个简单的goroutine leak检测库，可用于在测试运行结束后打印出所测试程序泄露的goroutine的stacktrace以及goroutine被创建的文件名和行号。
 7. 选择三个常见golang组件（channel, goroutine, [], map, sync.Map等），列举它们常见的严重伤害性能的anti-pattern。
 8. 一个C/C++程序需要调用一个go库，某一export了的go函数需高频率调用，且调用间隔需要调用根据go函数的返回调用其它C/C++函数处理，无法改变调用次序、相互依赖关系的前提下，如何最小化这样高频调用的性能损耗？
 9. 不考虑调度器修改本身，仅考虑runtime库的API扩展，如果要给调度器添加NUMA awareness且需要支持不同拓扑，runtime库需要添加哪些函数，哪些函数接口必须改动。
 10. stw的pause绝大多数情况下在100us量级，但有时跳增一个数量级。描述几种可能引起这一现象的触发因素和他们的解决方法。
 11. 已经对GC做了较充分优化的程序，在无法减小内存使用量的情况下，如何继续显著减低GC开销。
 12. 有一个常见说法是“我能用channel简单封装出一个类似sync.Pool功能的实现”。在多线程、high contention、管理不同资源的前提下，两者各方面性能上有哪些显著不同。
 13. 为何只有一个time.Sleep(time.Millisecond)循环的go程序CPU占用显著高于同类C/C++程序？或请详述只有一个goroutine的Go程序，每次time.Sleep(time.Millisecond)调用runtime所发生的事情。
 14. 一个Go程序如果尝试用fork()创建一个子进程，然后尝试在该子进程中对Go程序维护的一个数据结构串行化以后将串型化结果保存到磁盘。上述尝试会有什么结果？
 15. 请列举两种不同的观察GC占用CPU程度的方法，观察方法无需绝对精确，但需要可实际运用于profiling和optimization。
 16. GOMAXPROCS与性能的关系如何？在给定的硬件、程序上，如何设定GOMAXPROCS以期获得最佳性能？
 17. 一个成熟完备久经优化的Golang后台系统，程序只有一个进程，由golang实现。其核心处理的部分由数十个goroutine负责处理数万goroutine发起的请求。由于无法设定goroutine调度优先级，使得核心处理的那数十个goroutine往往无法及时被调度得到CPU，进而影响了处理的延迟。有什么改进办法？
 18. 列举几个近期版本runtime性能方面的regression。

 ### 高质量的 golang 面试题目
 1、https://www.zhangboxing.com/_posts/2019/02/01/golang%E9%9D%A2%E8%AF%95%E9%A2%98%E6%80%BB%E7%BB%93/
 2、https://www.zhihu.com/question/60952598
 
 
 
 