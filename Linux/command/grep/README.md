### 相关

### Op
1. 统计当前目录下 "t" 字符的出现次数 grep -c "t" -R -n .
n 代表打印行号， R 代表递归
2. grep "pattern" filename 
3. 打印匹配文本位于哪一个文件中。 grep -l linux sample.txt sample1.txt
4. 某个匹配结果之后的三行，seq 10 | grep 5 -A 3
5. 某个匹配结果之前的三行，使用 -B 选项。 seq 10 | grep 5 -B 3
6. 某个匹配结果之前以及之后的 3 行，使用 -C 选项。 seq 10 | grep 5 -C 3
7. 

