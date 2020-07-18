# 官方标准库的 token bucket 

## 本文内容
- 三种调用方式
- 源码的实现差异
- 理解 token bucket 与 leaky bucket 的差异

### 构造限流器 
```golang 
func NewLimiter(r Limit, b int) *Limiter {
	return &Limiter{
		limit: r,
		burst: b,
	}
}
```

- limit 表示xxx  burst 表示

(用如下图进行理解)


### 三种调用方式




### 与 leaky bucket 的对比
- 比 leaky bucket 能够应对突发流量（流量整形）trafic shaping