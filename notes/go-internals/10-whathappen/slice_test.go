package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	t.Log(cap(s1))
	t.Log(s1)
	Append(s1, 5)
	t.Log(s1)

	add(s1, 5)
	t.Log(s1)

	_ = append(s1, 5)
	t.Log(s1)
}

func Append(sl []int, num int) {
	sl = append(sl, num)
}

func add(s1 []int, addNum int) {
	for _, s := range s1{
		s += addNum
	}
}


func TestSliceAppend(t *testing.T) {
	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

// RESULT: 使用 append 向 slice 中添加元素，第一题中slice容量为5，所以补5个0，第二题为0，所以不需要。
//



func TestSliceAppend2(t *testing.T) {
	// a := []int{7, 8, 9}

	a := make([]int, 0, 10)
	a = append(a, 7, 8, 9)

	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)

	appSum(a)
	fmt.Printf("%+v\n", a)
}

// 因为append导致底层数组重新分配内存了，append中的a这个alice的底层数组和外面不是一个，并没有改变外面的。
func ap(a []int) {
	a = append(a, 10)
}

// 为什么，传的是引用
func app(a []int) {
	a[0] = 1
}

func appSum(a []int) {
	for idx, _ := range a {
		a[idx] = a[idx] * 5   // 会影响原数组
	}
}