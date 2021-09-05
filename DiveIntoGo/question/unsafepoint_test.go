package question

import (
	"fmt"
	"testing"
	"unsafe"
)

//- unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象，
// uintptr 类型的目标会被回收；
//- unsafe.Pointer 可以和 普通指针 进行相互转换；
//- unsafe.Pointer 可以和 uintptr 进行相互转换

type W struct {
	b int32
	c int64
}

func TestUnsafePointer(t *testing.T) {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b,w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b,w.c)
}