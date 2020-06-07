# 根据第一列，以逆序形式排序
sort -nrk 1 data.txt

# 依据第二列进行排序
sort -k 2 data.txt

sort -z data.txt | xargs -0

# 忽略不必要的字符
sort -bd data.txt


