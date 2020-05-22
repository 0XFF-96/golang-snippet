package preventLeak

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPreventLeak(t *testing.T) {
	doWork := func(strings <-chan string) <-chan interface{} {
		complete := make(chan interface{})
		go func() {
			defer close(complete)
			defer fmt.Println("doWork exist")
			for s := range strings {
				fmt.Println(s)
			}
		}()

		// 在 内部 goroutine 没有完成之前
		// 会阻塞在返回这里
		return complete
	}

	doWork(nil)

	// c := doWork(nil)
	// _, ok := <-c
	// 就会发生死锁🔒
	//fmt.Println(ok)
	fmt.Println("Done.")
}

func TestPreventLeakV3(t *testing.T) {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream exit") // 这一句永远不会执行...
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	newStream := newRandStream()
	fmt.Println("3 ints.")
	for i := 0; i < 3; i++ {
		fmt.Println(<-newStream)
	}
	// 这个会造成 goroutine 泄漏吗？
	// 没有任何东西证明..
}

func TestPreventLeakV4(t *testing.T) {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream exit")
			defer close(randStream)
			for {
				select {
				case <-done:
					return
				default:
					randStream <- rand.Int()
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	newStream := newRandStream(done)
	fmt.Println(" 3 ints")
	for i := 0; i < 3; i++ {
		fmt.Println(<-newStream)
	}
	// 可以这样来关
	close(done)
}

