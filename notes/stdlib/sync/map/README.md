### map 相关参考
1. https://developer.aliyun.com/article/741441
2. https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c

### 相关知识点
1. 这种转换的含义： *(*interface{})(p) 
2. atomic.StorePointer(&e.p, unsafe.Pointer(i))
3. atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i))
4. 需要弄清楚 unsafe.Pointer 的相关用法

gophercn: https://www.youtube.com/watch?v=C1EtfDnsdDs
相关的东西： https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c 

### 学习
1. 如何简单画出一个 package 的关系图/流程图？
2. 通过简单的代码，揭示内存布局和运行原理。

