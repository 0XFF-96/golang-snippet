一、关于 hookSignals 参数

- signal-drawin.go
- signal-window.go 
- signal-linux.go
这三个文件下，都有同一个函数 hookSignals 
是怎么通过，go env 绑定到当前的运行中的进程里面？


```
// initialize application
func (app *Application) initialize() {
	app.initOnce.Do(func() {
		app.servers = make([]server.Server, 0)
		app.workers = make([]worker.Worker, 0)
		app.logger = xlog.JupiterLogger
		app.signalHooker = hookSignals
		app.defers = []func() error{}
	})
}
```

二、关于启动流程顺序
- 看漏了这个
```golang
	err := xgo.SerialUntilError(fns...)()
```

