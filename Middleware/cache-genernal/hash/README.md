
### 什么是一致性哈希？
1. 一致性哈希图解：https://segmentfault.com/a/1190000021199728
2. 如何实现有权重的一致性哈希？
3. 为什么需要单独出来一个 hash func ?
4. 为什么用 entropy 验证？https://en.wikipedia.org/wiki/Entropy_(information_theory)
```
	mi := make(map[interface{}]int, len(keys))
	for k, v := range keys {
		mi[k] = v
	}
	entropy := mathx.CalcEntropy(mi)
```