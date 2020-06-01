### 可移植性
1、https://tonybai.com/2017/06/27/an-intro-about-go-portability/
2、https://blog.csdn.net/qq_36183935/article/details/81093323  
golang 深入剖析，初始化，编译器、目标文件，链接器 
3、https://tonybai.com/2005/09/28/also-talk-about-byte-order/ 什么是字节序？ 
大端、小端、网络字节序
4、什么是共享库？https://tonybai.com/2010/12/13/also-talk-about-shared-library/

### Go 可移植性

- 可以移植性的两点考量: 1、语言自身被移植到不同平台的容易程度 2、通过这种语言编译出来的应用对平台的适应性

#### 什么是 runtime ?
- runtime是支撑程序运行的基础
- 提供基础函数库调用
- 封装 syscall


#### cgo对可移植性的影响?

#### 什么是 internal linking 和 external linking ?


#### 如何实现纯静态连接？
在CGO_ENABLED=1这个默认值的情况下，是否可以实现纯静态连接呢？答案是可以。
在$GOROOT/cmd/cgo/doc.go中，文档介绍了cmd/link的两种工作模式：internal linking和external linking

### 命令学习

一、otool -L + target 命令
otool:  针对目标文件的展示工具，用来发现应用中使用到了哪些系统库，调用了其中哪些方法，使用了库中哪些对象及属性。
`otool -L hellogo`

相关介绍: https://www.cnblogs.com/sonofelice/p/5571416.html


二、nm target | grep " U " 命令
nm: 目标文件格式分析器。

