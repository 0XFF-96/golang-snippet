# 消息队列学习


### 疑问点
1、// chan chan 的类型是指什么？为什么这样用？
```Golang
type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}
```
2、dispatch、job、worker 三者之间的关系是什么？

3、

```任务分发的过程
				fmt.Printf("fetching workerJobQueue for: %s\n", job.Name)
				workerJobQueue := <-d.workerPool
				fmt.Printf("adding %s to workerJobQueue\n", job.Name)
				workerJobQueue <- job
```

## Reference 
1、https://gist.github.com/harlow/dbcd639cf8d396a2ab73
2、inspired by this article:http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
3、http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/ 【非常不错的文章】【】


## 消息队列设计
1、https://github.com/uw-labs/substrate
2、[用 redis 实现消费队列](https://juejin.im/entry/5a24e43ff265da432e5bd731)
3、