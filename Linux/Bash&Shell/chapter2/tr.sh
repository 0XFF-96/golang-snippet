

# 大小写转换
echo "HELOO WHO IS THIS" | tr 'A-Z' 'a-z'


# tr [option] set1 set2
# 输入字符从 set1 映射到 set2

echo 87654 | tr '987654321' '0-9'

# 加密
echo "tr name, tr saw tr conquered" | tr 'a-zA-Z' 'n-za-mN-ZA-M'

cat example.txt | tr -d '[set1]'

# set1 的补集
# 将补集的字符全部删除
#
echo hello 1 char 2 next 4 | tr -d -c '0-9 \n'


# 用 tr 压缩字符
# 压缩字符
