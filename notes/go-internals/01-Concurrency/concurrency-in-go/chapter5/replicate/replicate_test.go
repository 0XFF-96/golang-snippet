package replicate

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestReplicate_request(t *testing.T){
	doWork := func(
		done <-chan interface{},
		id int,
		wg *sync.WaitGroup,
		result chan<-int,
	){
		started := time.Now()
		defer wg.Done()

		// Simulate random load
		simulatedLoadTime := time.Duration(1 + rand.Intn(5)) *time.Second

		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		}

		select {
		case <-done:
		case result <-id:

		}

		took := time.Since(started)
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("%v took %v \n", id, took)
	}
	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i:=0;i <10;i++{
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	// 这个代表什么？
	//
	//
	fmt.Printf("Received an answer from %v \n", firstReturned)
}