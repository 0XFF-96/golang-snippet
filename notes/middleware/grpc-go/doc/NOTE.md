### 问题 
一、为什么 combine 的函数写得这么奇怪？
- 需要利用 copy 实现？
- 如何复现一下他里面的情况？overwritten element ?
它的意思是，o1 可能会有占位的情况，一些占位，可能会被写覆盖。
- 为什么会造成 data race ? slice 的传递
```
func combine(o1 []CallOption, o2 []CallOption) []CallOption {
	// we don't use append because o1 could have extra capacity whose
	// elements would be overwritten, which could cause inadvertent
	// sharing (and race conditions) between concurrent calls
```

