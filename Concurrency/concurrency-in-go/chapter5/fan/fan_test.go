package fan

import (
	"sync"
)

// One of the interesting properties of pipelines
// is the ability they give you to operate on the
// stream of data using a combination of separate,
// reorderable stages...
// 如何编写 fan-out
// 什么是 fan-in
var fanIn = func(
	done <-chan interface{},
	channels ...<-chan interface{}, ) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()

		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}
	// select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}