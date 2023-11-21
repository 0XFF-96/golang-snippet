// example4.go
package main

import "math/rand"

func fb2(s []int, index int) {
	_ = s[index:] // 第7行： 需要边界检查
	_ = s[:index] // 第8行： 边界检查消除了！
}

func fc2() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = s[:4]
	index := rand.Intn(7)

	// 位置互换了一下而已
	_ = s[index:] // 第15行： 需要边界检查
	_ = s[:index] // 第16行： 边界检查消除了！
}

func main() {}



