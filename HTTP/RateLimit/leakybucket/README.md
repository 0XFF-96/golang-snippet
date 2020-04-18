# this is the example code 

- i extracted from jaeger-client-go
- it implements a leaky bucket algorithm in a easy-understanding way.  
- this is the most understandable leaky bucket algorithm implemented in golang.


## what's the difference between leaky bucket and token bucket ?
- 漏桶算法能够强行限制数据的传输速率，令牌桶算法能够在限制数据的平均传输速率的同时还允许某种程度的突发传输


## Reference
- https://www.geeksforgeeks.org/leaky-bucket-algorithm/ [对比]
- https://en.wikipedia.org/wiki/Leaky_bucket 
- https://www.jianshu.com/p/a59c13e70582
- https://maiyang.me/post/2017-05-28-rate-limit-algorithm/ 
- https://www.slideshare.net/vimal25792/leaky-bucket-tocken-buckettraffic-shaping