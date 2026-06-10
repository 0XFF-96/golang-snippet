# 没有 -d 参数
echo "splitXsplitXsplitXsplit" | xargs -d X

# 读取 stdin， j
cat args.txt | xargs -n 1 bash echo.sh
