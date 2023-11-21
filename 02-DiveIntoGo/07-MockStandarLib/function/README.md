### internal 

1. https://github.com/teh-cmc/go-internals/blob/master/chapter1_assembly_primer/README.md
2. 这部分需要一点汇编的知识才能完全明白。
3. first class functions in go , https://go.dev/blog/functions-codewalk



### 汇编知识
1. 入参从右到左按顺序压栈并拷贝参数。但是参数的计算是从左到右

example 
```
package main

import "fmt"

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func foo(a int) int {
	fmt.Printf("%d\n", a)
	return a
}

func main() {
	myFunction(foo(66), foo(77))
}

// output:
// 66
// 77
```

// 指针传递的时候，复制的是指针的指针