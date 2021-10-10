package mock_interface

import (
	"fmt"
	"io"
	"os"
	"testing"
	"unsafe"
)

// From here
// https://www.infoq.cn/article/aLj4lqIgyNuLgmxHQrhq?utm_source=related_read_bottom&utm_medium=article
func TestM(t *testing.T) {
	var r io.Reader
	// {1}
	fmt.Printf("initial r: %T, %v\n", r, r)

	tty, _ := os.OpenFile("/Users/qcrao/Desktop/test", os.O_RDWR,1 )

	fmt.Printf("tty: %T, %v\n", tty, tty)

	// 给 r 赋值

	r = tty

	fmt.Printf("r: %T, %v\n", r, r)

	rIface := (*iface)(unsafe.Pointer(&r))

	fmt.Printf("r: iface.tab._type = %#x, iface.data = %#x\n", rIface.tab._type, rIface.data)

	// 给 w 赋值

	var w io.Writer

	w = r.(io.Writer)

	fmt.Printf("w: %T, %v\n", w, w)

	wIface := (*iface)(unsafe.Pointer(&w))

	fmt.Printf("w: iface.tab._type = %#x, iface.data = %#x\n", wIface.tab._type, wIface.data)

	// 给 empty 赋值

	var empty interface{}

	empty = w

	fmt.Printf("empty: %T, %v\n", empty, empty)

	emptyEface := (*eface)(unsafe.Pointer(&empty))

	fmt.Printf("empty: eface._type = %#x, eface.data = %#x\n", emptyEface._type, emptyEface.data)
}

// rtype must be kept in sync with ../runtime/type.go:/^type._type.
//《快学 Go 语言第十五课——反射》