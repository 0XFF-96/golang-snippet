


flag 库是 golang 提供命令行解析的一个标准库。

它有常见的三种用法：

一、绑定变量
```golang
		import "flag"
		var ip = flag.Int("flagname", 1234, "help message for flagname")
```

另一个种绑定变量的形式
```
		var flagvar int
		func init() {
			flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
		}
```


二、自定义命令行参数解析

你可以实现 `value` 接口的自定义命令行参数
```
you can create custom flags that satisfy the Value interface (with
	pointer receivers) and couple them to flag parsing by
		flag.Var(&flagVal, "name", "help message for flagname")
```


注意⚠️，在程序运行前，你需要调用 `flag.Parse()` 函数。
