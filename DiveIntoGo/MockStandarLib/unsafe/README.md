### Unsafe

1. unsafe 是一种可以绕过类型系统的限制，直接操作内存的方法
很危险⚠️，只有在特定场景下才能够使用。 

2. 有了解过 unsafe 包吗？

```
Package unsafe contains operations 
that step around the type safety of Go programs.
Packages that import unsafe may be non-portable 
and are not protected by the Go 1 compatibility guidelines
```

### 对 Go 指针对理解
1. Go 的指针不能进行数学运算
2. 不同类型的指针不能相互转换
3. 不同类型的指针不能使用 == 或 != 比较, 不能互相辅值

### case
1. 强制类型转换的例子： 
https://medium.com/a-journey-with-go/go-cast-vs-conversion-by-example-26e0ef3003f0

2. reflection 

```
func ValueOf(i interface{}) Value {
   [...]
   return unpackEface(i)
}
// unpackEface converts the empty interface i to a Value.
func unpackEface(i interface{}) Value {
   e := (*emptyInterface)(unsafe.Pointer(&i))
   [...]
}
```

3. The reflection also uses the unsafe package to modify 
the value of the reflected variable by updating the value directly in memory

### unsafe 内存布局
1. Go unsafe 包之内存布局, https://www.flysnow.org/2017/07/02/go-in-action-unsafe-memory-layout.html

### Reference 
1. https://medium.com/a-journey-with-go/go-what-is-the-unsafe-package-d2443da36350
unsafe 包有哪些应用？
Go: What is the Unsafe Package?
- Type safety, var k uint8 = *(*uint8)(unsafe.Pointer(&i))
- Usage in Go with the reflection package, 
- Usage in Go with the sync package
- Usage in Go with the runtime package, freeing stack memory、stack allocation
- 

2. https://www.cnblogs.com/qcrao-2018/p/10964692.html 【解密 unsafe 包】【五星】
3. Go: Cast vs Conversion by Example
4. https://go101.org/article/pointer.html 什么是 pointer ?