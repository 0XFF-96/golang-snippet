package main

import "fmt"

// FROM: https://www.cnblogs.com/qcrao-2018/p/10766091.html#3-iface-%E5%92%8C-eface-%E7%9A%84%E5%8C%BA%E5%88%AB%E6%98%AF%E4%BB%80%E4%B9%88
func main() {
	x := 200
	var any interface{} = x
	fmt.Println(any)

	g := Gopher{"Go"}
	var c coder = g
	fmt.Println(c)
}

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}