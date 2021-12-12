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
