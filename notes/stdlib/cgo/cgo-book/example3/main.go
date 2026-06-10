package main

//void SayHello(const char* s);
import "C"

func main(){
	C.SayHello(C.CString("Hello, World\n"))
}

// 报错了，未知原因
//# command-line-arguments
//Undefined symbols for architecture x86_64:
//"_SayHello", referenced from:
//__cgo_b1d487d8dfac_Cfunc_SayHello in _x002.o
//(maybe you meant: __cgo_b1d487d8dfac_Cfunc_SayHello)
//ld: symbol(s) not found for architecture x86_64
//clang: error: linker command failed with exit code 1 (use -v to see invocation)