package confinement

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func TestConfinement(t *testing.T) {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- i
		}
	}
	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}


func TestLexicalConfinement(t *testing.T) {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			// %c ？
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
		// 这个语句不会进行数据竞争
		// 因为 slice 复制进函数的
		// data = []byte("go123")
	}
	var wg sync.WaitGroup
	wg.Add(2)
	b := []byte("golang")

	// 这里传入的参数需要特别小心
	// 因为可能出出现数据修改的相关 bug
	go printData(&wg, b[:3])
	go printData(&wg, b[3:])

	fmt.Printf("%c", b)
	wg.Wait()
}