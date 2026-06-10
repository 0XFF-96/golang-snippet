package tipsAndTrick

import (
	"fmt"
	"testing"
)

type MyError string

func (e *MyError) Error() string {
	return string(*e)
}

var ErrBad = MyError("ErrBad")

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p // Will always return a non-nil error.
}

func returnsError2() error {
	return nil
}

func returnsError3() error {
	var p *MyError
	return p
}

// 涉及到  空接口、nil 的相关知识
// err是一个interface类型变量，其underlying有两部分组成：类型和值。只有这两部分都为nil时
// from here https://tonybai.com/2015/10/30/error-handling-in-go/
// https://golang.org/doc/faq#nil_error
func TestB(t *testing.T) {
	err := returnsError()
	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Println("return non-nil error")
		return
	}

	err = returnsError2()
	if err != nil {
		fmt.Println("error2 is not nil")
	}

	err = returnsError3()
	if err != nil {
		fmt.Println("error3 is not nil")
	}
	fmt.Println("return nil")
}