
### work pool 模式

目的：
- 控制 goroutine 的并发数量
- 让有缓存的 buffer chan hold 住并发量超过一定限制的任务。 

补充：
- 图->

```
type workerPool struct {
	concurrent int
	queues     []chan *Job
	jobPool    *sync.Pool
}
```

```
func NewWorkerPool(
	concurrent int,
	buffSize int,
) *PushWorker {
	w := &PushWorker{
		concurrent: concurrent,
		queues:     make([]chan *pushJob, 0, concurrent),
		jobPool: &sync.Pool{
			New: func() interface{} {
				return &Job{}
			},
		},
	}

	for i := 0; i < concurrent; i++ {
		ch := make(chan *pushJob, buffSize)
		w.queues = append(w.queues, ch)
		go w.worker(ch)
	}
	return w
}
```

```
func  worker(ch chan *pushJob) {
	for job := range ch {
		w.push(job)
	}
}
```

```
func (w *PushWorker) push() {
// real work to do here 
```


