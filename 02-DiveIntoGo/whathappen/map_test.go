package main

import (
	"fmt"
	"testing"
)

func TestTravelMap(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val   // 有问题
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// 注：key的顺序无法确定
// for range 循环的时候会创建每个元素的副本，而不是每个元素的引用，
// 所以 m[key] = &val 取的都是变量val的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，
// 因为最后 val 被赋值为3，所有输出的都是3
