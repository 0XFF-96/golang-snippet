package live_lock

import (
	"bytes"
	"fmt"
	"sync"
	. "sync/atomic"
	"testing"
	"time"
)

/ takeStep 的作用
// cadence 的作用
// 强调的重点： move at the same rate of speed
func TestLiveLock(t *testing.T){
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1*time.Microsecond){
			cadence.Broadcast()
		}
	}()
	takeStep := func(){
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprint(out, " ", dirName)
		AddInt32(dir, 1)
		takeStep()
		if LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()
		AddInt32(dir, -1)
		return false
	}

	var left, right int32

	tryLeft := func(out *bytes.Buffer) bool {return tryDir("left", &left, out)}
	tryRight := func(out *bytes.Buffer) bool {return tryDir("right", &right, out)}


	walk := func(walking *sync.WaitGroup, name string){
		var out bytes.Buffer

		defer func(){ fmt.Println(out.String())}()
		defer walking.Done()

		fmt.Fprint(&out, " is trying to scoot:", name)
		for i :=0;i < 5;i++{
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprint(&out, "\n tosses her hands up in exasperation!", name)
	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}

// 上面程序与原程序有出入
// success 了，可能有几个地方写错了。
// 1、This different-struct-memory demostrates a very common reason livelocks are written:
// two or more concurrent processes attemption to prevent a deadlock with coordination.
// 2、Livelocks are a subset of a larger set of problems called starvation.
//
