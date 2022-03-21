### 负载均衡

### 资料
1. 参考了《The power of two choices in randomized load balancing》的思路，
我们使用 the choice-of-2 算法，随机选取的两个节点进行打分，选择更优的节点:
(来自于 毛剑)

```
变更管理:
70％的问题是由变更引起的，恢复可用代码并不总是坏事。
避免过载:
过载保护、流量调度等。
依赖管理:
任何依赖都可能故障，做 chaos monkey testing，注入故障测试。
优雅降级:
有损服务，避免核心链路依赖故障。
重试退避:
退让算法，冻结时间，API retry detail 控制策略。
超时控制:
进程内 + 服务间 超时控制。
极限压测 + 故障演练。
扩容 + 重启 + 消除有害流量
```