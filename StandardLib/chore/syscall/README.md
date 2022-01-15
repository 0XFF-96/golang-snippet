### 从曹大的博客摘抄

> golang-notes 

可以把系统调用分为三类:

1. 阻塞系统调用
2. 非阻塞系统调用
3. wrapped 系统调用

提供给用户使用的系统调用，基本都会通知 runtime，以 entersyscall，exitsyscall 的形式来告诉 runtime，
在这个 syscall 阻塞的时候，由 runtime 判断是否把 P 腾出来给其它的 M 用。
解绑定指的是把 M 和 P 之间解绑，如果绑定被解除，在 syscall 返回时，这个 g 会被放入全局执行队列 globrunq 中。
同时 runtime 又保留了自己的特权，在执行自己的逻辑的时候，
我的 P 不会被调走，这样保证了在 Go 自己“底层”使用的这些 syscall 返回之后都能被立刻处理。
所以同样是 epollwait，runtime 用的是不能被别人打断的，
你用的 syscall.EpollWait 那显然是没有这种特权的。

## 参考资料

1. http://blog.studygolang.com/2016/06/go-syscall-intro/

2. the linux programming interface

3. https://mzh.io/golang-arm64-vdso

4. https://blog.csdn.net/luozhaotian/article/details/79609077