package question

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 空 slice 和 nil slice 之间有什么区别？
//
// nil切片和空切片指向的地址不一样。nil空切片引用数组指针地址为0（无指向任何实际地址）
// 空切片的引用数组指针地址是有的，且固定为一个值

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func TestEmptySlice(t *testing.T) {
	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}

// From: https://mp.weixin.qq.com/s/SHxcspmiKyPwPBbhfVxsGA
func TestForSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		s = append(s, v)
		fmt.Printf("len(s)=%v\n", len(s))
	}
}

// for range 源码
// The loop we generate:
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }