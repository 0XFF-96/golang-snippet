### Errors 案例分析

1. 用一个例子体会一下 `sql No rows error 的烦人之处`
2. 用一个优化的例子，看看上面的代码如何优化。 
3. 一眼看出问题所在， errors.WithStack(throwError()) 
4. 如何打印中间过程产生的参数？ 

- errors.WithMessagef(err, "param2:%d", param2), 利用 WithMessagef 貌似会把
WithStack 的调用栈给吞掉，导致又找不到原始的调用方法栈。徒劳无功...

分别观察一下下下面两种写法：
会产生两种不同的 log

```
log1:
// source=MiddleMethod2.func1 ？
// 信息解读: func1 是否指的是，这个函数内，第一个调用的函数？
// 这个信息是怎么被还原出来的？
2020/09/19 16:17:21 error is this a err message     error=param2:49: sql: no rows in result set source=MiddleMethod2.func1: /Users/lijinrui/golang-snippet/StandardLib/errors/error_better_example2/exmaple_test.go:80

log2:
2020/09/19 16:19:41 error is this a err message     error=param2:84: sql: no rows in result set

log3:
// 这种 log 我们能够知道，根源函数的情况，但是不知道中间过程是谁调用了他。 
2020/09/19 16:24:40 error is this a err message     error=sql: no rows in result set source=basicMethod2: /Users/lijinrui/golang-snippet/StandardLib/errors/error_better_example2/exmaple_test.go:124
```

（来对比一下这几个版本 log 的详细情况）

```golang
func MiddleMethod2()( err error ){
	param2 := rand.Intn(100)

    //---------
    // 区别的代码
	defer func()  {
		if err != nil {
			// err =  errors.WithMessagef(err, "param2:%d", param2)   // 1 
			err = errors.Wrapf(err, "param2:%d", param2)              // 2
		}
	}()
	//-------

	err = basicMethod()
	if err != nil {
		return err
	}

	err = basicMethod2()
	if err != nil {
		return err
	}

	err = basicMethod3()
	if err != nil {
		return err
	}
	return nil
}
```


Why ? 为什么呢？ 需要从实现原理上解析这种差异， 

到此为止，我们基本上能够解决一大半头痛的问题了。 
- 重复打印日志。
- 重复代码。
- 不知道根源函数和调用栈丢失的问题。

What's more ?

怎么返回错误给客户端？


5. 我们需要在(出口层)做 统一的 log 处理机制 ！！！！！


6. 最后，重点强调一下。 

```
//只附加新的信息
func WithMessage(err error, message string) error

//只附加调用堆栈信息
func WithStack(err error) error

//同时附加堆栈和信息
func Wrap(err error, message string) error
```


### 错误处理原则
1、errors 相关的东西。 
2、平时普通返回， 什么情况下需要用 wrap ?
3、什么情况需要用 wrap ? error 来自于外部的， 不能被本程序处理。 
4、当外部接口不够完善时，用 wrap 来包装更多信息。 
5、用【接口】来规范，错误行为。 
6、控制 error 的构造方法， 来规范 error 的分类。 
7、如何对于一个库，进行错误设计（custom error )。 如果针对调用者进行错误设计， 如何对于运维/devops 进行错误哦设计。

