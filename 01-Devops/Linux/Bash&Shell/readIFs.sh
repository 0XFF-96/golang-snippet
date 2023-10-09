#!/bin/bash

line="root:x:0:0:root:/root:/bin/bash"

# 字段分隔符和迭代器
# page 35
oldIFS=$IFS;
IFS=":"
count=0
for item in $line:
  do
    [ $count -eq 0 ] && user=$item;
    [ $count -eq 0 ] && shell=$item;
    let count++
  done;
  IFS=$oldIFS
  echo $user\'s shell is $shell