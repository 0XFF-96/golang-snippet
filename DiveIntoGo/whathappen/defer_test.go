package main

import (
	"fmt"
	"testing"
)

func TestDefercall(t *testing.T) {
	defer_call()
}

func defer_call()  {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()

	panic("触发异常")
}

// RESULT
// NOTE: 出现 panic 语句的时候，会先按照 defer 的后进先出顺序执行


// FROM: https://goplay.tools/snippet/0z1WdNUY5eX
func TestDeferMethod(t *testing.T) {
	testmethod()
	testmethod2()
	testmethod3()

	testmethod4_value()
	testmethod4_pointer()
}

func testmethod() {
	var i = 1
	defer func() {
		var c = func() int {
			return i * 2
		}()
		fmt.Println("defer", c)
	}()
	i = i + 1
}

func testmethod2() {
	var i = 1
	defer func() {
		c := i * 2
		fmt.Println("defer", c)
	}()
	i = i + 1
}

func testmethod3() {
	var i = 1
	defer func() {
		fmt.Println("defer", i*2)
	}()
	i = i + 1
}

func testmethod4_value() {
	var i = 1
	defer func(i int) {
		var c = func() int {
			return i * 2
		}()
		fmt.Println("4_v defer", c)
	}(i)
	i = i + 1
}

func testmethod4_pointer() {
	var i = 1
	defer func(i *int) {
		var c = func() int {
			return *i * 2
		}()
		fmt.Println("4_p defer", c)
	}(&i)
	i = i + 1
}
