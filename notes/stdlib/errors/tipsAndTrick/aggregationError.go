package tipsAndTrick

import (
	"fmt"
)

// 聚合error handle functions

//err := doStuff1()
//if err != nil {
////handle A
////handle B
//... ...
//}
//
//err = doStuff2()
//if err != nil {
////handle A
////handle B
//... ...
//}
//
//err = doStuff3()
//if err != nil {
////handle A
////handle B
//... ...
//}

func handleA() {
	fmt.Println("handle A")
}
func handleB() {
	fmt.Println("handle B")
}

func doStuff1()error{return nil }
func doStuff2() error {return nil}
func doStuff3() error {return nil }

func foo() {
	var err error
	defer func() {
		if err != nil {
			handleA()
			handleB()
		}
	}()

	err = doStuff1()
	if err != nil {
		return
	}

	err = doStuff2()
	if err != nil {
		return
	}

	err = doStuff3()
	if err != nil {
		return
	}
}