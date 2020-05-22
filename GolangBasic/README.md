### Processing
- context 包（还可以从哪些地方进行挖掘？）
- flag 包



### 规律总结
总结一下看 【阅读 go 源码库的相关规律】
1、为什么需要阅读这个库？这个库，在你日常生活中使用频率有多高？阅读完这个库之后，能带给你怎么样的收益？
2、基本使用
3、学会看测试用例，这些测试用例分别是测试哪些逻辑？
4、从 Go 源码中，git checkout 一下这部分的代码？看看经过了哪些方面的变迁？有没有重大 bug 
5、你从这个库学到了什么？ 你在哪里看到其他地方有使用这个库？ 这个库的难点？


### Golang 源码阅读列表
（如何进行记录） ？
drwxr-xr-x    8 lijinrui  staff    256 Dec 16 14:11 bufio
drwxr-xr-x    4 lijinrui  staff    128 Dec 16 14:11 archive
drwxr-xr-x    3 lijinrui  staff     96 Dec 16 14:11 builtin
drwxr-xr-x   12 lijinrui  staff    384 Dec 16 14:11 bytes
drwxr-xr-x   27 lijinrui  staff    864 Dec 16 14:11 cmd
drwxr-xr-x    8 lijinrui  staff    256 Dec 16 14:11 compress
drwxr-xr-x    5 lijinrui  staff    160 Dec 16 14:11 container
drwxr-xr-x    8 lijinrui  staff    256 May  4 21:57 context
drwxr-xr-x   23 lijinrui  staff    736 Dec 16 14:11 crypto
drwxr-xr-x    3 lijinrui  staff     96 Dec 16 14:11 database
drwxr-xr-x    8 lijinrui  staff    256 Dec 16 14:11 debug
drwxr-xr-x   14 lijinrui  staff    448 Dec 16 14:11 encoding
drwxr-xr-x    7 lijinrui  staff    224 Dec 16 14:11 errors
drwxr-xr-x    4 lijinrui  staff    128 Dec 16 14:11 expvar
drwxr-xr-x    7 lijinrui  staff    224 Dec 16 14:11 flag
drwxr-xr-x   15 lijinrui  staff    480 Dec 16 14:11 fmt
drwxr-xr-x   14 lijinrui  staff    448 Dec 16 14:11 go
drwxr-xr-x   11 lijinrui  staff    352 Dec 16 14:11 hash
drwxr-xr-x    9 lijinrui  staff    288 Dec 16 14:11 html
drwxr-xr-x   19 lijinrui  staff    608 Dec 16 14:11 image
drwxr-xr-x    3 lijinrui  staff     96 Dec 16 14:11 index
drwxr-xr-x   22 lijinrui  staff    704 Dec 16 14:11 internal
drwxr-xr-x   10 lijinrui  staff    320 Dec 16 14:11 io
drwxr-xr-x    6 lijinrui  staff    192 Dec 16 14:11 log
-rwxr--r--    1 lijinrui  staff   6844 Dec 16 14:11 make.bash
-rw-r--r--    1 lijinrui  staff   4007 Dec 16 14:11 make.bat
drwxr-xr-x  132 lijinrui  staff   4224 Dec 16 14:11 math
drwxr-xr-x   19 lijinrui  staff    608 Dec 16 14:11 mime
drwxr-xr-x  194 lijinrui  staff   6208 Dec 16 14:11 net
drwxr-xr-x  102 lijinrui  staff   3264 Dec 16 14:11 os
drwxr-xr-x    8 lijinrui  staff    256 Dec 16 14:11 path
drwxr-xr-x    6 lijinrui  staff    192 Dec 16 14:11 plugin
-rwxr--r--    1 lijinrui  staff   1021 Dec 16 14:11 race.bash
-rw-r--r--    1 lijinrui  staff   1041 Dec 16 14:11 race.bat
drwxr-xr-x   22 lijinrui  staff    704 Dec 16 14:11 reflect
drwxr-xr-x   14 lijinrui  staff    448 Dec 16 14:11 regexp
-rwxr--r--    1 lijinrui  staff   2155 Dec 16 14:11 run.bash
-rw-r--r--    1 lijinrui  staff   1164 Dec 16 14:11 run.bat
-rwxr--r--    1 lijinrui  staff    435 Dec 16 14:11 run.rc
drwxr-xr-x  536 lijinrui  staff  17152 Dec 16 14:11 runtime
drwxr-xr-x   19 lijinrui  staff    608 Dec 16 14:11 sort
drwxr-xr-x   26 lijinrui  staff    832 Dec 16 14:11 strconv
drwxr-xr-x   16 lijinrui  staff    512 Dec 16 14:11 strings
drwxr-xr-x   25 lijinrui  staff    800 Dec 16 14:11 sync
drwxr-xr-x  316 lijinrui  staff  10112 Dec 16 14:11 syscall
drwxr-xr-x    3 lijinrui  staff     96 Dec 16 14:11 testdata
drwxr-xr-x   21 lijinrui  staff    672 Dec 16 14:11 testing
drwxr-xr-x    5 lijinrui  staff    160 Dec 16 14:11 text
drwxr-xr-x   32 lijinrui  staff   1024 Dec 16 14:11 time
drwxr-xr-x   14 lijinrui  staff    448 Dec 16 14:11 unicode
drwxr-xr-x    3 lijinrui  staff     96 Dec 16 14:11 unsafe
drwxr-xr-x    4 lijinrui  staff    128 Dec 16 14:11 vendor


### 相关函数 