
#!/bin/bash

BACKUPFILE=backup-$(date +%m-%d-%Y)
# 在备份文件中嵌入时间.
archive=${1:-$BACKUPFILE}
#  如果在命令行中没有指定备份文件的文件名,
#  那么将默认使用"backup-MM-DD-YYYY.tar.gz".

tar cvf - `find . -mtime -1 -type f -print` > $archive.tar
gzip $archive.tar
echo "Directory $PWD backed up in archive file \"$archive.tar.gz\"."

exit 0