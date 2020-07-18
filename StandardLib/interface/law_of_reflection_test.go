package _interface

import (
	"fmt"
	"reflect"
	"testing"
)

// 1. Reflection goes from interface value to reflection object.
// both Type and Value have a Kind method that returns a constant indicating
// what sort of item is stored: Uint, Float64, Slice, and so on.
func TestTypeOf(t *testing.T) {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	// -----------------
	var y float64 = 3.4
	v := reflect.ValueOf(y)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// -----
	var z uint8 = 'x'
	v1 := reflect.ValueOf(z)
	fmt.Println("type:", v1.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v1.Kind() == reflect.Uint8) // true.
	z = uint8(v1.Uint())
	// v.Uint returns a uint64.

	// 性质2:
	// ---
	// The second property is that the Kind of a reflection object describes the underlying type,
	// not the static type. If a reflection object contains a value of a user-defined integer type, as in
	type MyInt int
	var k MyInt = 7
	v2 := reflect.ValueOf(k)
	fmt.Println(v2)
	fmt.Println(v2.Type())          // MyInt
	fmt.Println(reflect.TypeOf(k))  // MyInt
	fmt.Println(v2.Kind())          // int
	fmt.Println(reflect.TypeOf(v2)) // reflect.Value
}

// 2. Reflection goes from reflection object to interface value.
// 不太懂文章这里例子想表达些什么...
func TestInterface(t *testing.T) {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println(v)

	// panic
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)
}

// 3. To modify a reflection object, the value must be [settable].
// 什么是 settable ?
// 请与其他语言的进行区别!
//

func TestSettable(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1) // Error: will panic.

	// why ?
	//  it's that v is not settable.
	//  Settability is a property of a reflection Value,
	//  and not all reflection Values have it.
	fmt.Println(v.CanSet())

	// CanSet 的判断逻辑是
	// it's that v is not settable.
	// Settability is a property of a reflection Value,
	// and not all reflection Values have it.
	//

	// Setability
	// 1、Settability is determined by whether the reflection object holds the original item
	// 2、

	// var x float64 = 3.4
	// v1 := reflect.ValueOf(x1)
	// v1.SetFloat(7.1)
	// 如果允许更新，会造成什么后果？
	// f(x)
	// f(&x)
	// 上面例子不能更新的原因，和 f(x), f(&x) 不能更新的原因是一样的。
}

func TestSetability(t *testing.T) {
	var x1 float64 = 3.4
	p := reflect.ValueOf(&x1) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	// it's not p we want to set, it's (in effect) *p.
	//  it represents x
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x1)
}

// when using reflection to modify the fields of a structure.
// As long as we have the address of the structure, we can modify its fields.
//

type T struct {
	A int
	B string
}

func TestStruct(t *testing.T) {
	tt := T{23, "skidoo"}
	s := reflect.ValueOf(&tt).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

// 看看这个小例子，
// 关于一个 error 的 interface 的例子。
func TestInr(t *testing.T) {
	var i interface{}
	// i := &interface{
	i = 1
	t.Log("i.(type)")
	if i != nil {
		t.Log("i")
		switch i.(type) {
		case int:
			t.Log("int")
		}
	} else {
		t.Log("not nil")
	}
}
