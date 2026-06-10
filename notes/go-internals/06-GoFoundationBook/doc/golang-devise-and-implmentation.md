
# go 语言设计与实现（读书笔记）

1、https://draveness.me/golang/   

一、 map 主题。  
1、[runtime.mapaccess2](https://draveness.me/golang/tree/runtime.mapaccess2) ，
2 两个函数的区别。 字面量的意思 2、初始化、增、删、改、查、扩容。等操作对应的函数，已经影响。 
2、[runtime.mapassign](https://draveness.me/golang/tree/runtime.mapassign)  
map[k] 赋值入口、扩容的入口是 [runtime.hashGrow](https://draveness.me/golang/tree/runtime.hashGrow) 
 4、[runtime.mapdelete](https://draveness.me/golang/tree/runtime.mapdelete)    

二、函数调用
略。   

三、接口
1、接口的本质， 接口是计算机系统中多个组件共享的边界，
不同的组件能够在边界上交换信息[1](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/#fn:1)。
如下图所示，接口的本质是引入一个新的中间层，调用方可以通过接口与具体实现分离，
解除上下游的耦合，上层的模块不再需要依赖下层的具体模块，只需要依赖一个约定好的接口。   
2、接口， 类型转换、类型断言、动态派发机制 等。
3、go 语言的接口实现是隐式的，在【编译期间】对代码进行类型检查。 
4、数据结构 runtime.iface、runtime.eface 
5、itab.fun, 它是一个用于动态派发的虚函数表，存储了一组函数指针。
虽然该变量被声明成大小固定的数组，但是在使用时会通过原始指针获取其中的数据，
所以 fun 数组中保存的元素数量是不确定的 。  
6、动态派发， 动态派发（Dynamic dispatch）是在运行期间选择具体【多态操作】（方法或者函数）执行的过程，它是面向对象语言中的常见特性。 

四、https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/
