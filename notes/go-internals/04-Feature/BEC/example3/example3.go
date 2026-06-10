
// example3.go
package main

import "math/rand"

func fa() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	index := rand.Intn(7)
	_ = s[:index] // 第9行： 需要边界检查
	_ = s[index:] // 第10行： 边界检查消除了！
}

func fb(s []int, index int) {
	_ = s[:index] // 第14行： 需要边界检查
	_ = s[index:] // 第15行： 需要边界检查（不够智能？）
}

func fc() {
	s := []int{0, 1, 2, 3, 4, 5, 6}

	// 多了这个..
	s = s[:4]
	index := rand.Intn(7)
	_ = s[:index] // 第22行： 需要边界检查
	_ = s[index:] // 第23行： 需要边界检查（不够智能？）
}

func main() {}


// 知识盲点
// 为什么呢？ 原因是一个子切片表达式中的起始下标可能会大于基础切片的长度。 让我们先看一个简单的使用了子切片的例子：
//

// _ = s0[:index]
// 上一行是安全的不能保证下一行是无需边界检查的。
// 事实上，下一行将产生一个恐慌，因为起始下标
// index大于终止下标（即切片s0的长度）。
// _ = s0[index:] // panic