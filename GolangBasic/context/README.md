### TODO 

1、继续整理 Context 代码官方文章
2、写重点函数解析文章
3、git checkout 一下官方的库，看看有什么东西值得查看。


### 文章
1、https://blog.xizhibei.me/2019/08/26/golang-the-highly-controversial-context/ 
2、https://learnku.com/articles/29877 
3、https://learnku.com/articles/29877 
4、https://zhuanlan.zhihu.com/p/34417106 
5、幕布上之前整理过的文章集合：https://mubu.com/doc3c36FqaM2_ 

### 官方 Context Eample 最佳实践
1- 用法1: 用于防止 goroutine leak 泄露
2- 设置 cancel 方法，用于主动删除场景
3- 设置超时时间
4- 用于上下文传递 value 。k-v 值的形式（jaeger 中深有体会）

### UnitTest 测试例子

- ParentFinishesChild、
- TestDeadline、TestTimeOut、CanceledTimeout(当我们需要编写与时间有关的函数时，可以借鉴这里的测试方法) 
例如，我们想要一个函数，每天 9 点执行一次。我们应该怎么写测试用例呢？你需要模拟一下 9 点吗？ 
又或者，我想要 A 函数调用 B 函数后的 1 分钟后，延迟执行 C 函数。 你应该如何为 A 函数编写单元测试？
- XTestValues, 使用不同 WithValue() 值的时候，看看是否会互相覆盖的情况。
- XTestAllocs: 这个是在测试些什么？ 看起来是在测试 GC 时候，Context 包的一些表现
```golang
// Short reports whether the -test.short flag is set.
func Short() bool {
```

```
// AllocsPerRun returns the average number of allocations during calls to f.
// Although the return value has type float64, it will always be an integral value.
//
// To compute the number of allocations, the function will first be run once as
// a warm-up. The average number of allocations over the specified number of
// runs will then be measured and returned.
//
// AllocsPerRun sets GOMAXPROCS to 1 during its measurement and will restore
// it before returning.
func AllocsPerRun(runs int, f func()) (avg float64) {
```

- TestSimultaneousCancels： 测试同时取消函数
- CancelRemoves: 测试在 child 因为 timeout 的情况下，parent 数据结构的改变
> 主要测试的是以下代码的逻辑：`func (c *cancelCtx) cancel(removeFromParent bool, err error) {` 