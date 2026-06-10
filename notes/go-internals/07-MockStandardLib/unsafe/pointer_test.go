package unsafe

import (
	"fmt"
	"reflect"
	"testing"
)

// 什么是 pointers ?
// https://go101.org/article/pointer.html
// 1. Memory Addresses
// 2. Value Addresses

// QA
// How to Get a Pointer Value and What Are Addressable Values?
// new(T) or Use & operator to take the address of t.
// Generally speaking, an addressable value means a value which is hosted at somewhere in memory.

// Go Pointer Types and Values 四种情况：
//*int  // A non-defined pointer type whose base type is int.
//**int // A non-defined pointer type whose base type is *int.
//// Ptr is a defined pointer type whose base type is int.
//type Ptr *int
//// PP is a defined pointer type whose base type is Ptr.
//type PP *Ptr


// 限制
// 1. Go pointer values don't support arithmetic operations
// 2. A pointer value can't be converted to an arbitrary pointer type
// 3. A pointer value can't be compared with values of an arbitrary pointer type
// 4. A pointer value can't be assigned to pointer values of other pointer types

//  address taking and pointer dereference
func TestAddress(t *testing.T) {
	p0 := new(int)   // p0 points to a zero int value.
	fmt.Println(p0)  // (a hex address string)
	fmt.Println(*p0) // 0

	// x is a copy of the value at
	// the address stored in p0.
	x := *p0
	b := &x
	fmt.Println("x:", x, &x)
	*b = 10
	fmt.Println(x, b)

	fmt.Println("----part2-----")
	// Both take the address of x.
	// x, *p1 and *p2 represent the same value.
	p1, p2 := &x, &x
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false, why ?  p0 -> int, p1 -> x
	fmt.Println(reflect.TypeOf(p0), reflect.TypeOf(p1))
	fmt.Println("------------")

	p3 := &*p0 // <=> p3 := &(*p0) <=> p3 := p0
	// Now, p3 and p0 store the same address.
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}


func double(x *int) {
	*x += *x
	x = nil // the line is just for explanation purpose
}

// 下面输出的结果
func TestExample1(t *testing.T) {
	var a = 3
	double(&a)
	fmt.Println(a) // 6
	p := &a
	double(p)
	fmt.Println(a, p == nil, p, &a) // 12 false
}


// Return Pointers of Local Variables Is Safe in Go
// 这会导致逃逸问题？
func newInt() *int {
	a := 3
	return &a
}