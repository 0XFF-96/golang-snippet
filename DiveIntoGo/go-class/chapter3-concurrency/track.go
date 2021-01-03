package chapter3_concurrency

import (
	"sync"
)

type Track struct {
	wg sync.WaitGroup
}

func (t *Track) Event(data []byte){

	t.wg.Add(1)
	go func() {
		t.wg.Done()
	}()
}

func (t *Track) ShutDown() {
	t.wg.Wait()
}

func (t *Track) ShutDownV2() {
}

// Question()
// 1. 如何使用 Track