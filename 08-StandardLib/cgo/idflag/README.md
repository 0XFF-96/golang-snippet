
### 流程
1、flags="-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git describe --long --dirty --abbrev=14`"
  go build -ldflags "$flags" -x -o build-version main.go
  
2、Git Commit Hash: 
  UTC Build Time : 2020-06-07_10:15:26AM


### symbol not found 的解决方法
```
local ldflags="-X ${PACKAGE}/util/define.Release=${COMMITID} -linkmode 'external' -extldflags '-static'"
```

### 为什么需要知道编译平台的信息？
很有必要。尤其是我们发布二进制的时候，必须清楚说明你自己的编译平台。
当然，这也不仅仅是针对 go，其他静态语言 c、c++ 也一样。
编译器编译的时候会做很多优化，所谓的指令重排、内存逃逸等等。
这些都可能有潜在的 bug，一些知名项目 k8s、codis 发布二进制的时候也都基本会注入编译平台的版本号


### 其他人怎么用 idflag ?

etcd-io/etcd/functional/build
8:CGO_ENABLED=0 go build -v -installsuffix cgo -ldflags "-s" -o ./bin/etcd-agent ./functional/cmd/etcd-agent

