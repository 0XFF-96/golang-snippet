### 什么是 free 命令

我们通过free命令查看机器空闲内存时，会发现free的值很小。这主要是因为，在Linux系统中有这么一种思想，内存不用白不用，因此它尽可能的cache和buffer一些数据，以方便下次使用。但实际上这些内存也是可以立刻拿来使用的。 

###  知识点
- 数据来源：从/proc/meminfo中读出的出来的数值
- 


### 命令
- 释放掉被系统cache占用的数据: echo 3>/proc/sys/vm/drop_caches 
- 



### 官方文档阅读
1. https://www.computerhope.com/unix/free.htm
2. https://man7.org/linux/man-pages/man1/free.1.html

3. 重要 https://www.cnblogs.com/coldplayerest/archive/2010/02/20/1669949.html
(提供了几个计算公式)

4. https://blog.csdn.net/loongshawn/article/details/51758116
5. https://www.cnblogs.com/ggjucheng/archive/2012/01/08/2316438.html