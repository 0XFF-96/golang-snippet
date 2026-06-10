# GMP 模型


### 1.1 GMP 模型版本

```
static void
schedule(G *gp)
{
 ...
 schedlock();
 if(gp != nil) {
  ...
  switch(gp->status){
  case Grunnable:
  case Gdead:
   // Shouldn't have been running!
   runtime·throw("bad gp->status in sched");
  case Grunning:
   gp->status = Grunnable;
   gput(gp);
   break;
  }

 gp = nextgandunlock();
 gp->readyonstop = 0;
 gp->status = Grunning;
 m->curg = gp;
 gp->m = m;
 ...
 runtime·gogo(&gp->sched, 0);
}
```
1. 来自 https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&mid=2247487503&idx=1&sn=bfc20f81a1c6059ca489733b31a2c63c&scene=21#wechat_redirect


### GO 1.14 调度器发生很大的变化

1、单核 CPU，开两个 Goroutine，其中一个死循环，会怎么样？
2、代码来源:  https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&mid=2247487643&idx=1&sn=f81b18a12ab156feebb9fc9329e1c8f4&scene=21#wechat_redirect
```
// Main Goroutine 
func main() {
    // 模拟单核 CPU
    runtime.GOMAXPROCS(1)
    
    // 模拟 Goroutine 死循环
    go func() {
        for {
        }
    }()

    time.Sleep(time.Millisecond)
    fmt.Println("脑子进煎鱼了")
}
```

### go 如何启动？
1、详解 Go 程序的启动流程，你知道 g0，m0 是什么吗？


