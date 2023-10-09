### TODO
1、nsq 的 make file 文件写的不太好，没有提供一个统一的命令，能够从源码启动各个组件
所有，可以考虑用 docker 的形式，起一个 all-in-one 的 demo nsq应用。或者 后台启动程序。 

2、


### 值得学习的地方

一、flag set 命令
```
func nsqdFlagSet(opts *nsqd.Options) *flag.FlagSet {
	flagSet := flag.NewFlagSet("nsqd", flag.ExitOnError)
```


二、version 的作用
```
	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String("nsqd"))
		os.Exit(0)
	}
```

三、利用 go-svc 库，抽象服务启动过程

```
To get started, implement the Init, Start, and Stop methods to do
any work needed during these steps.

1、Init and Start cannot block. Launch long-running your code in a new Goroutine.
2、Stop may block for a short amount of time to attempt clean shutdown.
3、Call svc.Run() with a reference to your svc.Service implementation to start your program.

结论：When running in console mode Ctrl+C is treated like a Stop Service signal.
For a full guide visit https://github.com/judwhite/go-svc
```

四、nsqd 真正的启动过，做了哪些事情？

```
func New(opts *Options) (*NSQD, error) {
	var err error
```


五、

```
	err = p.nsqd.LoadMetadata()
	err = p.nsqd.PersistMetadata()
```


六、在这里实现用 decorate 的设计模式

- 优缺点？
- 有什么作用？
```
func Decorate(f APIHandler, ds ...Decorator) httprouter.Handle {
	decorated := f
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		decorated(w, req, ps)
	}
}
```

Reference:
1、https://tech.meituan.com/2016/07/01/mq-design.html
2、https://studygolang.com/articles/25313 [nsq 源码剖析]
3、https://nsq.io/overview/design.html [设计文档]