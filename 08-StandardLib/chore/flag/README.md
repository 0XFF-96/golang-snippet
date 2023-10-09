### Flag 官方有三种用法

### Flag 标准库阅读
一、Example 1: A single string flag called "species" with default value "gopher".
二、// Example 2: Two flags sharing a variable, so we can have a shorthand.
  // The order of initialization is undefined, so make sure both use the
  // same default value. They must be set up with an init function.
三、Example 3: A user-defined flag type, a slice of durations.
四、flag-set, 在 nsq or jaeger repo 里面有看见这种用户，借鉴一下，用模板复制一下

### 学习资料
1. 看看 douyu 的 flag 怎么写...  pkg/flag 模块的实现
2. 
3. https://supereagle.github.io/2017/07/01/golang-flag/ 
4. 实现一个简单的 flag, https://zhuanlan.zhihu.com/p/136572155
5. https://supereagle.github.io/2017/07/01/golang-flag/

### 疑惑🤔 
一、能否解析如下参数？
go run main.go   --fix=1,2,3,4 -hi=1:2,3:4 


### 学习笔记
核心数据结构： flagSet 和 flag 
核心函数： VisitAll 和 Parse (parseone)


一、值得学习的错误设计处理方式
```
		switch f.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
```

