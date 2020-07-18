package main

// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import "fmt"

// Calling C function pointers is currently not supported,
// however you can declare Go variables which hold C function pointers
// and pass them back and forth between Go and C.
// C code may call function pointers received from Go
func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 42
}
