package tipsAndTrick

import (
	"fmt"
	"testing"
)

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func Foo() error {
	return MyError("foo error")
}

func TestF(t *testing.T) {
	err := Foo()
	switch e := err.(type) {
	default:
		fmt.Println("default")
	case error:
		fmt.Println("found an error:", e)
	case MyError:
		fmt.Println("found MyError:", e)
	}
	return
}

func TestF2(t *testing.T) {
	err := Foo()
	switch e := err.(type) {
	default:
		fmt.Println("default")
	case error:
		fmt.Println("found an error:", e)
	case MyError:
		fmt.Println("found MyError:", e)
	}
	return
}
