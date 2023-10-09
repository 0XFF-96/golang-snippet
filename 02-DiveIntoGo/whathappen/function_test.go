package main

import (
	"fmt"
	"testing"
)

func hello() []string {
	return nil
}

func Testfun(t *testing.T) {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

