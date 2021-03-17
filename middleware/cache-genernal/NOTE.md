### 目标
	.	缓存数据的存储;
	.	过期数据项管理;
	.	内存数据导出/导入;
	.	提供CRUD接口


5、https://my.oschina.net/slagga/blog/4480099
6、实现缓存系统， https://zhuanlan.zhihu.com/p/81867557 。 。 
经典库之一： https://github.com/patrickmn/go-cache 。 
经典库之二： https://github.com/allegro/bigcache 。 
经典库之三：https://github.com/golang/groupcache 。 
经典库之四：https://pandaychen.github.io/2020/07/02/A-GOLANG-CCACHE-ANALYSIS/ 。 cccache 。 
https://pandaychen.github.io/2020/07/02/A-GOLANG-CCACHE-ANALYSIS/。 
其他：https://blog.csdn.net/weixin_33519829/article/details/112098752
相关： https://talks.golang.org/。 

7、go-zero 缓存 https://mp.weixin.qq.com/s/ueo9zNMx1OO4cScebORN9Q 
8、https://coolshell.cn/articles/17416.html

9、singleflight 算法，主要是用于做缓存击穿相关的东西的。  
关于测试
1、如何测试这个 cache 库的指标？
2、如何知道是否有内存泄漏？
3、如何进行性能测试？
4、GC 的效率如何？
5、成本不高的优化点在哪里？
