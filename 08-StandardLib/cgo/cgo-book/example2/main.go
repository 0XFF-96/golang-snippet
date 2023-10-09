package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/
import "C"
// 需要把代码放在 import "C" 的上面才行。

func main(){
	C.SayHello(C.CString("Hello, World\n"))
}