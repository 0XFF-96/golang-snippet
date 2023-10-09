### 介绍
1. 这是 MIT 6.824 分布式课程的一个作业 map reduce
2. 


### 参考文章
1、http://marcio.io/2015/07/cheap-mapreduce-in-go/



### 为什么单独执行单测试，会出错？

```
GOROOT=/Users/lijinrui/go/go1.13.5 #gosetup
GOPATH=/Users/lijinrui/go #gosetup
/Users/lijinrui/go/go1.13.5/bin/go test -c -o /private/var/folders/09/r6f6jypj0f9bj7rz2qpv6jyr0000gn/T/___TestBasic_in_test_test_go__1_ -gcflags "all=-N -l" /Users/lijinrui/golang-snippet/Pattern/MapReduce-All/mapreduce/test_test.go #gosetup
# command-line-arguments [command-line-arguments.test]
```

- 用命令行 go test . -v 执行不会？
- go test -n . 可以看到 go test 被执行编译的整个过程