### big-cache 设计重点
1. set、get、delete 
2. 为什么 big-cache 需要把 value 存在字节数组里面？效率、方法？

``
1、queue.BytesQueue是一个字节数组，可以做到按需分配。当加入一个[]byte时，它会把数据copy到尾部

问题一：只有 startIdx , 如何定位 queue.BytesQueue 在字节数组中的 value ？
value = queue[startIdx, endIdx]
1. valueLen 的信息被 encode 到 header 里面了

key: index 
map[uint]int
``


### Reference 
1. https://colobu.com/2019/11/18/how-is-the-bigcache-is-fast/
2. https://blog.csdn.net/weixin_33519829/article/details/112098752
这篇文章对 big-cache 的内存布局有比较大的帮助，
尤其是如何实现 bytes queue 和 ring-buffer 方面。
3.  