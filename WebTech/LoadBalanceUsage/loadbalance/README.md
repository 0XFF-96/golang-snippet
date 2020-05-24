this is the code from this article:
https://levelup.gitconnected.com/create-a-simple-load-balancer-with-go-4b474460bab2

###  文章架构
实现了一个 Round Robin Selection 的 lb 算法



### 这篇文章有不少技术点值得学习

* Traverse the slice as a cycle。循环遍历 slice 
* We just want to increase the counter by one, we don't want to lock。利用 atomic 的特性（？）
* Avoid Race Conditions。Rmutex 适用于读多，写少的场景。
* Route traffic only to healthy backends，two ways: active、passive。
* Retry 逻辑