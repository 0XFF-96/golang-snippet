package repeat

import (
	"fmt"
	"math/rand"
	"testing"
)


var toString = func(
	done <-chan interface{},
	valueStream <-chan interface{},
) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

var take = func(
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
			case takeStream <- <-valueStream:

			}
		}
	}()
	return takeStream
}


func TestRepeat(t *testing.T) {
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

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{}, ) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}

var repeat = func(
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

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N), ) {
	}
}