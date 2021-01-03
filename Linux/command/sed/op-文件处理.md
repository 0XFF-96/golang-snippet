### SED 做文件处理

1、在每一行的前面和后面加上一些字符和数据。 
2、
参考方法 ： http://jianshu.com/p/a79955ff6374 
https://wangchujiang.com/linux-command/c/sed.html


- 去除重复行

sort file |uniq

- 查找非重复行

sort file |uniq -u

- 查找重复行
sort file |uniq -d

- 统计
sort file | uniq -c


### 在每一行前后添加字符

https://www.jianshu.com/p/a79955ff6374

1. 在每一行首 

2. 在每一行末尾

sed 's/$/&TAIL/g' test.file
sed 's/^/HEAD&/g' test.file


sed '/./{s/^//"&/;s/$/&/",/}' ids


sed 's/^/"&/g' ids2