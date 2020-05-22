package avalible_same

import (
	"fmt"
	"sync"
	"testing"
)

func TestIsVariableInSameSpace(t *testing.T){
	var wg sync.WaitGroup
	saluation := "hello"

	say := func(){
		defer wg.Done()
		saluation = "welcome"
	}
	wg.Add(1)
	go say()
	wg.Wait()
	fmt.Println(saluation)
}

func TestIsVAriableInSameSpace2(t *testing.T){
	saluation := []string{"hello", "welcom", "you"}
	var wg sync.WaitGroup

	for _, s := range saluation {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println(s)
		}()
		// 在这里放 wg.Wait() 才是正确的？
		// wg.Wait()
	}
	wg.Wait()
}

// 引用的地址, 你能想到的一切都与这个有关。
// 1、. Because the goroutines being scheduled may run at any point in time in the future,
// it is undetermined what values will be printed from within the goroutine.
// 2、salutation is transferred to the heap holding a reference to the last value in my string slice

// 和原书的实验结果有出入...
// 在这两个地方，wait 有什么区别？
func TestIsVAriableInSameSpace3(t *testing.T){
	saluation := []string{"hello", "welcom", "you"}
	var wg sync.WaitGroup

	for _, s := range saluation {
		wg.Add(1)
		go func(str string) {
			// 不能在 goroutine 内 wg.Add
			// wg.Add(1)
			defer wg.Done()
			fmt.Println(str)
		}(s)
		// 1
		// wg.Wait()
	}
	// 2
	wg.Wait()
	// 画出 1、2 的执行流程图
}