### select-inside-select 的设计
一、解析一下，这种设计的好处？
```
		select {
		case ch <- res:
		default:
			select {
			case <-ch:
			default:
			}
			ch <- res
		}
	}
```


### Channel 总结
1. https://segmentfault.com/a/1190000017958702
2. https://lessisbetter.site/about/  关于 channel 的相关源码
3. https://colobu.com/2016/04/14/Golang-Channels/
4. https://dave.cheney.net/2014/03/19/channel-axioms

### Close Chan 的用法之一 ctx.Done()
1. 来源: https://colobu.com/2016/04/14/Golang-Channels/ 

### Example 
1. 并发交替打印

### pipelines 模式
1. https://blog.golang.org/pipelines 

### 利用 context 控制 goroutine 的启动和终止 
1. https://blog.csdn.net/neweastsun/article/details/108244101
2. 
### Channel 的用处
把channel用在数据流动的地方：

1、消息传递、消息过滤
2、信号广播
3、事件订阅与广播
4、请求、响应转发
5、任务分发
6、结果汇总
7、并发控制
8、同步与异步