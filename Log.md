## Log

### Repo changelog

- bilibili Discovery 的使用
- grpc 的 resolver 和 balancer 之间有什么关系？和 register 有什么关系？
- middleware/cache: 缓存 cache / lru 算法
- Sync map 和 standardLib 的 map ，以及他们对于实现 go-zero 的 timewheel 的相关影响
- 抽象一下 middleware/cache-general 的相关方法
- 增加 snippet 文件目录，收集日常觉得惊艳的代码片段
- 斗鱼开源框架 juno
- middleware/cache-general/bigcache 内存缓存框架分析
- StandardLib/binary 编码的原理和实现分析 varint。什么是 x080 ?
- 错误处理原则
- middleware/grpc-go 的相关源码阅读
- StandardLib/recover, recover 原理
- StandLib/Escape 逃逸分析
- middleware/cache-general: asyncache 库 from bytedance
- pkg/AccessControl grafana 权限控制和管理的源码阅读
- pkg/process -> go-zero/proc 统一的 signal 和 graceful stop down 管理器
- pkg/notify -> grafana notify 系统设计
- DiveIntoGo/MockStandardLib: mock 一遍基础库的相关代码
- FrameWork/k8s client go 源码阅读
- 增加 operation 目录，用于记录一些运维操作 Or 部署一些组件的操作
- 内存 Cron job 的框架 https://github.com/robfig/cron
- 欧神的个人代码仓库: https://github.com/changkun/pkg
- 用 go 语言实现的推荐系统框架 https://github.com/zhenghaoz/gorse
- 增加 etcd/raft 的相关实现
- 重新整理整个项目的相关结构
- 优化 Goland + ideaVim 的工作使用流程，提高一些敲代码的效率
- 重新调整了本项目的结构，把 Go 源码学习的目录再整理了一下
- 增加 Terraform 的项目

### Reading index

1. http://xiaorui.cc/archives/7172 [readlog]
2. https://segmentfault.com/a/1190000020086816 — 为什么 resp.Body.Close() 需要 close ?
3. golang 热门技术: https://github.com/aceld/golang
4. 刷题模板: https://github.com/greyireland/algorithm-pattern
5. 各个公司架构: https://github.com/davideuler/architecture.of.internet-product
6. https://www.ifanr.com/1108702 — 最好的产品，永远是下一款！针对当前社交 APP 的相关分析
7. 纯文本管理 OKR: https://www.bmpi.dev/self/life-in-plain-text/
8. 增加 database/mysql 源码解析的相关内容。Note: http、rpc、redigo、go-sql-driver 是开发一个服务常见的四大组件
9. golang 优秀文章集合: https://github.com/ardanlabs/gotraining/blob/master/reading/README.md
10. https://github.com/ffhelicopter/Go42 — 开源 go 书籍📚，对常用框架有总结
11. https://www.zhihu.com/question/314214128/answer/681836873 — 如何解决 IM 消息中的时序问题
12. grpc deadline 问题: https://xiaomi-info.github.io/2019/12/30/grpc-deadline/
13. golang 内存模型 (Russ Cox): https://research.swtch.com/mm
14. http://arthurchiao.art/blog/rbac-as-it-meant-to-be-zh/ [五🌟] — RBAC: 演进历史、设计理念及简洁实现 (Tailscale, 2021)
15. tech interview handbook: https://github.com/yangshun/tech-interview-handbook
16. mysql 事务机制的实现: https://www.cnblogs.com/rickiyang/p/13652664.html
17. golang 文档风格的注释: https://github.com/fluhus/godoc-tricks
18. golang 满血补给包: https://github.com/0voice/Introduction-to-Golang
19. 如何阅读源码？https://www.codedump.info/post/20200605-how-to-read-code-v2020/
20. go-micro 连接池实现: https://github.com/asim/go-micro/blob/94bd1025a64eaaba1d89189891ef49f3ab5eee7f/util/pool/default.go

### 关于产品

- 认知偏差知识手册 [from bytedance]: https://s75w5y7vut.feishu.cn/docs/doccn3BatnScBJe7wD7K3S5poFf
