### 一、什么是 idflags ?
1、idflags 在编译器中的两个作用。 
2、http://blog.champbay.com/2019/11/26/ldflags%E5%9C%A8golang%E7%BC%96%E8%AF%91%E4%B8%AD%E7%9A%842%E4%B8%AA%E4%BD%9C%E7%94%A8/
3、在【链接】 过程起相应作用。

### 动态链接的缺点：
一个可执行应用，在A服务器上进行了编译–>链接操作，在链接的时候所依赖的那些函数库是存在A服务器上的。
当把这个可执行应用拿到B服务器上去运行，因为在B服务器上该可执行应用所依赖的函数库可能缺少或者位置不正确，而导致报错。
所以，在使用这个可执行应用之前，需要在B服务器上进行make一下生成所对应的可执行应用才可以。
这也就是 configure–>make–>make install的真正含义所在

### 二、在二进制包中动态注入，版本信息
1. https://ms2008.github.io/2018/10/08/golang-build-version/ 

### 三、inject build time variabel 
1.https://blog.alexellis.io/inject-build-time-vars-golang/

### 四、golang在编译时用ldflags设置变量的值
1. https://studygolang.com/articles/9422

### 五、 隐藏二进制信息
1. https://studygolang.com/articles/7496


Thank you for this tip! I've been beating my head against the wall on this for a while.

### tony.bai 的文章，关于 go 语言的移植性，感觉不错。
1. https://tonybai.com/2017/06/27/an-intro-about-go-portability/

CGO_ENABLED=1的情况下，也编译构建出了一个纯静态链接的Go程序。 

### external linek 