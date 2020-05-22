package generator

import (
	"fmt"
	"testing"
)

// As a reminder, a
// generator for a pipeline is any function that
// converts a set of discrete values into a
// stream of values on a channel.
func TestHandyGenerators(t *testing.T) {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			defer fmt.Println("repeat exists")
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			// defer close(done)
			defer fmt.Println("take exists!")
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				// 这个符号挺特别的..
				// 为什么需要两次 <- <- ?
				// 有可能会有什么问题？
				//
				case takeStream <- <-valueStream:

				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	// 这一条和我在 take 里面设置的那一条 close 相比如何？
	// xxx
	// count := 0
	r := repeat(done, 10, 9)
	for num := range take(done, r, 10) {
		//count++
		//if count == 5 {
		//	close(done)
		//}
		fmt.Printf("%v ", num)
	}

	t.Log("-------")
	for num := range take(done, r, 5) {
		//count++
		//if count == 5 {
		//	close(done)
		//}
		fmt.Printf("%v ", num)
	}
	t.Log("------")
}