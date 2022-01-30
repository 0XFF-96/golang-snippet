
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


### Worker pool 模式 

- From https://www.kancloud.cn/aceld/golang/2043313
 
```
type WorkerManager struct {
   //用来监控Worker是否已经死亡的缓冲Channel
   workerChan chan *worker
   // 一共要监控的worker数量
   nWorkers int
}

//创建一个WorkerManager对象
func NewWorkerManager(nworkers int) *WorkerManager {
   return &WorkerManager{
      nWorkers:nworkers,
      workerChan: make(chan *worker, nworkers),
   }
}

//启动worker池，并为每个Worker分配一个ID，让每个Worker进行工作
func (wm *WorkerManager)StartWorkerPool() {
   //开启一定数量的Worker
   for i := 0; i < wm.nWorkers; i++ {
      i := i
      wk := &worker{id: i}
      go wk.work(wm.workerChan)
   }

  //启动保活监控
   wm.KeepLiveWorkers()
}

//保活监控workers
func (wm *WorkerManager) KeepLiveWorkers() {
   //如果有worker已经死亡 workChan会得到具体死亡的worker然后 打出异常，然后重启
   for wk := range wm.workerChan {
      // log the error
      fmt.Printf("Worker %d stopped with err: [%v] \n", wk.id, wk.err)
      // reset err
      wk.err = nil
      // 当前这个wk已经死亡了，需要重新启动他的业务
      go wk.work(wm.workerChan)
   }
}
```