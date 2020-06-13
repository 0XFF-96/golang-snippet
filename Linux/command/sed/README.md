### 使用 sed 进行文本替换
1. sed 会将每一行中符合模式的内存替换掉，但是如果要替换
所有内容，我们需要在命令尾部加上参数 g 
2. 

### op 
1. sed 's/pattern/replace_string' file
2. cat file | sed 's/pattern/replace_string' file
3. echo thisthisthisthiTHISTHISTHIS | sed 's/this/THIS/2g'
只需要从 第 2 个匹配开始替换
4. 注意定界符的问题
sed 's|te\|xt|replace|g'

5. 移除空白行 sed '/^$/d' file 
6. 直接在文件中进行替换
sed 's/PATTERN/replacement' -i filename

sed -i 's/\b[0-9]\{3\}\b/NUMBER/g' sed_data.txt


7. sed 中，利用双引号会通过对表达式求数值进行扩张，想在 sed 表达式中使用一些变量时，就能
派上用场。




 
