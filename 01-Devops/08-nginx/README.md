### Nginx 相关命令

netstat -n | awk '/^tcp/ {++state [$NF]} END {for (key in state) print key,"t",state [key]}'


