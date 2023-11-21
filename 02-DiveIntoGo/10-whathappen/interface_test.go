package main

import (
	"fmt"
	"testing"
)

type T interface {}

func TestInterfaceEqual(t *testing.T) {
	var tt T
	var tt2 interface{}
	var x interface{} = nil
	var tt3 T = nil
	t.Log(tt == tt2)
	t.Log(tt == x)
	t.Log(tt == tt3)
	t.Log(tt3 == x)


	t.Log(tt)
	t.Log(tt2)
	// 忘记它怎么写的了
	// 初始化空接口和判定的问题。
	//

	// golang 空接口的相关问题
}


// 当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil
func TestInter(t *testing.T) {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	i = nil
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("i is not nil")
	}

	var err error
	var i2 interface{}
	i2 = err
	if i2 == nil {
		fmt.Println("i2 is nil")
	} else {
		fmt.Println("i2 is not nil")
	}
}
