### 重读经典文章

1、https://blog.golang.org/laws-of-reflection 入门必读系列 Rob pike 写的 3-15
https://segmentfault.com/a/1190000006190038 6-22

2、https://research.swtch.com/interfaces [未看]
3、https://juejin.im/post/5a6873fd518825734501b3c5 [未看，对接口对深入理解]
4、http://legendtkl.com/2017/06/12/understanding-golang-interface/
5、https://www.infoq.cn/article/s9EJ0VLVKP8sSIQYi4X8?utm_source=related_read_bottom&utm_medium=article 
- 反射三定律
- 反射的应用场景
- 反射需要注意的事项（未导出的类型） 

--==== 未看 ---===
1、https://juejin.im/post/5d31884af265da1baa1eae79
2、https://i6448038.github.io/2018/10/01/Golang-interface/ 底层数据结构

4、https://juejin.im/post/5eeab6c3f265da02e12b9daa [深入浅出了解 interface 和 go 反射的关系]
5、interface 的类型与反射三定律 https://segmentfault.com/a/1190000006190038 [done]
6、https://juejin.im/post/6844904177009688589 

4、interface 的一个视频，很好说清楚了类型
https://www.youtube.com/watch?v=F4wUrj6pmSI&t=123s
6、https://www.cnblogs.com/qcrao-2018/p/10822655.html#%E5%8F%8D%E5%B0%84%E7%9B%B8%E5%85%B3%E5%87%BD%E6%95%B0%E7%9A%84%E4%BD%BF%E7%94%A8

7、
// 从 golang-notes
// 中 copy 过来的一个文档
// 结合这篇文章，将一下 reflect 和 interface
// 还有，go 语言圣经中的文章。

8、https://www.infoq.cn/article/s9EJ0VLVKP8sSIQYi4X8?utm_source=related_read_bottom&utm_medium=article

### 反射的应用
1、IDE 中的代码自动补全功能? 
2、对象序列化（json 函数库）
3、fmt 相关函数的实现、
4、ORM（全称是：Object Relational Mapping，对象关系映射
5、reflection 与 DeepEqual 函数

### 反射的知识点
- 基本数据结构
- interface 的相关抽象 <type, value> 
- golang 的反射建立在类型系统的基础上。

### 结合 sync.Map 的相关知识点
1. https://stackoverflow.com/questions/38043678/golang-embedded-interface-on-parent-struct/38058463
2. https://stackoverflow.com/questions/29971363/golang-reflection-on-interface-vs-pointer-to-interface

### 结构体内嵌 interface 的用法[视频]
1、https://mubu.com/app/edit/home/2IHkQXa2J_ 


### 为什么 reflect 会造成性能下降？
1. 对象的反射生成和获取都会增加额外的代码指令，
并且也会涉及interface{}装箱/拆箱操作，中间还可能增加临时对象的生成，
所以性能下降是肯定的，但是具体能下降多少呢，还是得数据来说话。