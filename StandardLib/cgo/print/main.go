package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// 1、Conversion between Go and C strings is done with the C.CString, C.GoString, and C.GoStringN functions.
func Print(s string) {
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))

	//  需要 free memoery 的原因
	//  you must remember to free the memory when you're done with it by calling C.free.
	C.free(unsafe.Pointer(cs))
}

func PrintV2(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}


// 看看这里奇怪的 build 信息💻
// /Users/lijinrui/go/go1.13.5/bin/go build -o /private/var/folders/09/r6f6jypj0f9bj7rz2qpv6jyr0000gn/T/___go_build_main_go__1_ /Users/lijinrui/golang-Snippet/StandardLib/cgo/print/mapreduce.go.bak #gosetup
func main(){
	Print("234")

}