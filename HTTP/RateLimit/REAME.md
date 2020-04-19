# RateLimit 

- go-kit 服务限流：https://www.hwholiday.com/2020/go_kit_v4/
- https://blog.csdn.net/micl200110041/article/details/82013032
- 标准库限流：https://www.cyhone.com/articles/analisys-of-golang-rate/
- kong-网关：https://konghq.com/blog/how-to-design-a-scalable-rate-limiting-algorithm/ 
- leaky-bucket 的改进: https://www.cyhone.com/articles/analysis-of-uber-go-ratelimit/


## 仓库
    - uber 的 leaky-bucket 的实现：https://github.com/uber-go/ratelimit/
    - 官方 token bucket 的注释版 https://github.com/chenyahui/AnnotatedCode/tree/master/go/x/time
    - pkg/ratelimit, 我们自己利用 x/time/rate 官方库的能力，构建的 grpc method 级别的 ratelimte 
    