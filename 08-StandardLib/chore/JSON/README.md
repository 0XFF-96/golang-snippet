### 参考文章

1、json 标准库真的慢吗？for Benchmark
https://medium.com/a-journey-with-go/go-is-the-encoding-json-package-really-slow-62b64d54b148

    -? 在哪一步，loop on the feild and write byte in the buffer?
    -关键优化 ： https://github.com/golang/go/commit/c0547476f342665514904cf2581a62135d2366c3#diff-e79d4db81e8544657cb631be813f89b4
        在这 PR 中，利用 sync.Pool 优化了内存的使用。
    -go run -gcflags="-m, 在运行时，进行内存分析。

### Json 源码分析

### FQA
- `json:"xxx"`, 实现这个需要利用到反射， 具体的源码在哪里？
- 如何进行结构体递归处理？
- 如何避免♻️循环递归引用？
- json 中的 type encoder 和 value encoder 之间的作用？

- 真正执行代码的地方
```
func (e *encodeState) reflectValue(v reflect.Value, opts encOpts) {
	valueEncoder(v)(e, v, opts)
}
```