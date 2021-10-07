package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	// a := unsafe.Sizeof(unsafe.ArbitraryType(1))
	// t.Log(a)
}

type W struct {
	b int32
	c int64
}

//- unsafe.Pointer 只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算。
// 而 uintptr 是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象，

// uintptr 类型的目标会被回收；
//- unsafe.Pointer 可以和 普通指针 进行相互转换；
//- unsafe.Pointer 可以和 uintptr 进行相互转换

// QA: 什么是指针运算？
// QA: 类型如何被回收？
func TestUnsafePtrAndPinter(t *testing.T) {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b,w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))

	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b,w.c)
}

func TestTransfer(t *testing.T) {
	a :="aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v",b)
}

// 上面例句
// 那么如果想要在底层转换二者，只需要把 StringHeader 的地址强转成 SliceHeader 就行
// 那么go有个很强的包叫 unsafe 。
// 1.unsafe.Pointer(&a)方法可以得到变量a的地址
// 2.(*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式
// 3.(*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切片的指针
// 4.再通过 *转为指针指向的实际内容

func TestSliceLen(t *testing.T) {
	s := make([]int, 9, 20)

	// 为什么是 + uintprt(8) ?

	// runtime/slice.go
	//type slice struct {
	//	array unsafe.Pointer // 元素指针
	//	len   int // 长度
	//	cap   int // 容量
	//}

	// 转换流程如下
	// uintptr -> unsafe.Pointer -> *T
	// Len: &s => pointer => uintptr => pointer => *int => int
	// Cap: &s => pointer => uintptr => pointer => *int => int
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20
}

func TestMap(t *testing.T) {
	// 获取 map 的长度
	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}

// 应用三
// 在 map 源码中，mapaccess1、mapassign、mapdelete 函数中，需要定位 key 的位置，会先对 key 做哈希运算。
// 这个语句在干什么？
// 属于文档中提到的哪一个例子？
// b := (*bmap)(unsafe.Pointer(uintptr(h.buckets) + (hash&m)*uintptr(t.bucketsize)))

// 解构出来
// b := 5.(*bmap)(    4.unsafe.Pointer(   1.uintptr(h.buckets) + 2.(hash&m) * 3.uintptr(t.bucketsize)      ))
// 总共有 5 部分
// 最后解构出来的解构是 b
//


// map 源码
// store new key/value at insert position
//if t.indirectkey {
//	kmem := newobject(t.key)
//	*(*unsafe.Pointer)(insertk) = kmem
//	insertk = kmem
//}
//if t.indirectvalue {
//	vmem := newobject(t.elem)
//	*(*unsafe.Pointer)(val) = vmem
//}
//
//typedmemmove(t.key, insertk, key)

type Programmer struct {
	name string
	language string
}

func TestShift(t *testing.T) {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)
}


// 什么是 zero-copy
// 路径：src/reflect/value.go。只需要共享底层 []byte 数组就可以实现 zero-copy
//