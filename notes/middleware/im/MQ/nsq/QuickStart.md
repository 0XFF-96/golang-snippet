


### 第一步：从源码编译组件


### 启动 nsqd, 玩耍
1、nsqlookupd
2、./nsqd --lookupd-tcp-address=127.0.0.1:4160
3、./nsqadmin --lookupd-http-address=127.0.0.1:4161
4、./nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161

一定要按照上面顺序来启动，否则会造成 push message 的时候有， topic not found 的现象


## 组件介绍

### nsqlookupd

`nsqlookupd` is the daemon that manages topology metadata and serves client requests to
discover the location of topics at runtime.

Read the [docs](http://nsq.io/components/nsqlookupd.html)


## nsqd

`nsqd` is the daemon that receives, queues, and delivers messages to clients.

Read the [docs](http://nsq.io/components/nsqd.html)

## nsqadmin

`nsqadmin` is a Web UI to view aggregated cluster stats in realtime and perform various
administrative tasks.

Read the [docs](http://nsq.io/components/nsqadmin.html)

## Working Locally

 0. install nodejs 8.x (includes npm)
 1. `$ npm install`
 2. `$ ./gulp --series clean watch` or `$ ./gulp --series clean build`
 3. `$ go-bindata --debug --pkg=nsqadmin --prefix=static/build static/build/...`
 4. `$ go build ../apps/nsqadmin && ./nsqadmin`
 5. make changes (repeat step 4 only if you make changes to any Go code)
 6. `$ go-bindata --pkg=nsqadmin --prefix=static/build static/build/...`
 7. commit other changes and `bindata.go`


### 解读日志信息





### 数据流 Data flow 




### 