package starvation

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// http://c.biancheng.net/view/4786.html
func TestStavation(t *testing.T){
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second


	// 大小锁问题
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime;{
			sharedLock.Lock()
			time.Sleep(3 *time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %d work loops", count)
	}

	policteWorker := func(){
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) < runtime;{
			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %d work loops", count)
	}

	wg.Add(2)
	go greedyWorker()
	go policteWorker()

	wg.Wait()
}

// Keep in mind that starvation can also
// apply to CPU, memory, file handles, database connections:
// any resource that must be shared is a candidate for starvation

