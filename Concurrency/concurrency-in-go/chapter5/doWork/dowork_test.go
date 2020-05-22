package doWork

import (
	"testing"
	"time"
)

func DoWork(
	done <-chan interface{},
	nums ... int,
)(<-chan interface{}, <-chan int){
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)

	go func(){
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 *time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case intStream <- n:

			}
		}
	}()
	return heartbeat, intStream
}

func TestDoWork_GeneratesAllNumbers(t *testing.T){
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for i, expected := range intSlice {
		select {
		case r := <- results:
			if r != expected {
				t.Errorf("index %v: expected %v, but received %v",
					i,
					expected,
					r,
				)
			}
		case <-time.After(1 *time.Second):
			t.Fatal("test timed out")
		}
	}
}

func TestDoWork_GeneratestAllNumbers(t *testing.T){
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 3, 5}
	heartbeat, results := DoWork(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i];r != expected {
			t.Errorf("index %v expected %v but received %v",
				i,
				expected,
				r,
			)
		}
	}
}

// utilizing interval-based hearbeats !
