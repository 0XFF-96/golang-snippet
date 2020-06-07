### 知识点
- porting-go 
- print 为什么需要 free memory ?


### 实现资料
1、https://tonybai.com/2017/06/27/an-intro-about-go-portability/ [done]
2、经典 Cgo 的实验文章：https://blog.golang.org/cgo [done]
3、https://golang.org/cmd/cgo/    [done]
4、dave cheney 的相关文章：https://dave.cheney.net/tag/cgo


### 书籍📚
1、https://book.douban.com/subject/6801697/  C语言接口与实现


### cgo documentation 阅读
1、Using cgo with the go command 
2、Go references to C 
3、Passing pointers 
4、Special cases 
5、Using cgo directly 
结论：
1、
2、Calling variadic C functions is not supported. It is possible to circumvent this by using a C function wrapper.
3、C references to GO, 不太懂。因为，export、extern 两个关键词。 
