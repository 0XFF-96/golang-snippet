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



