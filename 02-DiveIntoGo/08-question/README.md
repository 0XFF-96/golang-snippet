### 目录&大纲

### 此目录收集一下 golang 常见的坑

1. 为什么我们需要及时关闭 ```resp.Close()``` ?
2. 空 slice 和 nil slice 之间有什么区别？
3. 能说说 uintptr 和 unsafe.Pointer 的区别吗？

https://mp.weixin.qq.com/s?__biz=MzAwMDAxNjU4Mg==&mid=2247483679&idx=1&sn=7075859e59741b1d0a81dc472b8ce45f&scene=21#wechat_redirect

4. 当连接一个不存在的 ip 时，发生了什么事情？

5. 简单聊聊内存逃逸？

https://mp.weixin.qq.com/s?__biz=MzAwMDAxNjU4Mg==&mid=2247483686&idx=1&sn=e48c51107191f02da5751a19a54f7d41&scene=21#wechat_redirect

### Chan 和 Select 
// 怎么用图表，把这些状态分类？
// 读、写、初始化、未初始化、已经关闭
// 有缓冲、无缓冲、缓冲满
// chanel 总结，https://colobu.com/2016/04/14/Golang-Channels/

1. for循环select时，如果通道已经关闭会怎么样？如果select中的case只有一个，又会怎么样？
2. 怎么样才能不读关闭后通道 ?
3. select里只有一个已经关闭的case，置为nil，会怎么样？
```
所以select里最好有一个default，否则将有一直阻塞的风险
```

4. 对未初始化对 chan 进行读写，会怎么样？为什么
5. 对已经关闭对 chan 进行读写，会怎么样？