package mutex

import (
	"fmt"
	"math"
	"os"
	"sync"
	"testing"
	"text/tabwriter"
	"time"
)

// 出现了这个错误
// panic: sync: negative WaitGroup counter, 有个地方 done 多了一次
// 这是对比 Rmutex 和 mutex 之间的性能差异相关的
// 问题 1
// 这个程序的临界区在哪里？
// criticial section ?
// where is the criticial section?
func TestRWMutex(t *testing.T){
	producer := func(wg *sync.WaitGroup, l sync.Locker){
		defer wg.Done()
		for i:=5; i > 0; i--{
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker){
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker)time.Duration{
		var wg sync.WaitGroup
		wg.Add(count+1)
		go producer(&wg, mutex)

		begineTime := time.Now()
		for i:=count;i>0;i--{
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(begineTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ',0)
	defer tw.Flush()

	fmt.Fprint(tw, "Readers\tRWMutext\tMutex\n")
	var m sync.RWMutex

	for i:=0;i<20;i++{
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}