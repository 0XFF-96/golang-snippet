package main

import (
	"fmt"
)

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

// go tool compile -S -N -l main.go 编译

// 为什么需要带 n 和 l 参数？
// 如果编译时不使用 -N -l 参数，编译器会对汇编代码进行优化，编译结果会有较大差别

func main() {
	myFunction(66, 77)


	// 验证 go 参数传递的基本方式
	// go 语言在传递参数时是【传值】还是【传引用】也是一个有趣的问题，
	// 不同的选择会影响我们在函数中修改入参时是否会影响调用方看到的数据
	i := 30
	arr := [2]int{66, 77}
	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	myFunction2(i, arr)
	fmt.Printf("after  calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

// Go 语言的整型和数组类型都是值传递的，也就是在调用函数时会对内容进行拷贝。
// 需要注意的是如果当前数组的大小非常的大，这种传值的方式会对性能造成比较大的影响
func myFunction2(i int, arr [2]int) {
	fmt.Printf("in my_funciton - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

// 思考🤔
// 1、根据 main 函数生成的汇编指令，我们可以分析出
// 2、（结论）
// 通过分析 Go 语言编译后的汇编指令，我们发现 Go 语言使用栈传递参数和接收返回值，所以它只需要在栈上多分配一些内存就可以返回多个值


// 我们通过指针修改结构体中的成员变量，结构体在内存中是一片连续的空间，指向结构体的指针也是指向这个结构体的首地址。
// 将 MyStruct 指针修改成 int 类型的，那么访问新指针就会返回整型变量 i，将指针移动 8 个字节之后就能获取下一个成员变量 j
//```
//type MyStruct struct {
//	i int
//	j int
//}
//
//func myFunction(ms *MyStruct) *MyStruct {
//	return ms
//}
//
//$ go tool compile -S -N -l main.go
//"".myFunction STEXT nosplit size=20 args=0x10 locals=0x0
//	0x0000 00000 (main.go:8)	MOVQ	$0, "".~r1+16(SP) // 初始化返回值
//	0x0009 00009 (main.go:9)	MOVQ	"".ms+8(SP), AX   // 复制引用
//	0x000e 00014 (main.go:9)	MOVQ	AX, "".~r1+16(SP) // 返回引用
//	0x0013 00019 (main.go:9)	RET
//```