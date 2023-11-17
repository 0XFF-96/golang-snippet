### order by 是如何工作？

1、可能使用内存完成，也可能需要外部排序
2、sort_buffer 是在 server 层的。 sort_buffer_size ,  
mysql 为排序开辟的内容 （sort_buffer) 的相关大小。 
3、rowid 排序算法 和 全字段排序算法， max_length_for_sort_data 这个参数进行控制。 
4、如果 Mysql 认为内存足够大，会有限选择全字段排序，把需要的字段都放到 sort_buffer 中， 
这样排序后就会直接从内存里面返回查询结果，不同再回到原表去取数据。 
5、Using filesort, 是需要内存中排序的意思，会对性能有一定程度影响。 
6、覆盖所有，指索引上的信息足够满足查询请求，不需要再回到主键索引上去取数据。  
7、Extra 字段里面多了 “Using index”， 表示的就是使用了覆盖索引，性能上会快很多。 