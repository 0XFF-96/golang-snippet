package context_switch

import (
	"sync"
	"testing"
)

// begin 的作用要重点看一下
// 为什么可以配合 close() 这样使用？
func BenchmarkContextSwitch(b *testing.B){
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	sender := func(){
		defer wg.Done()
		<-begin
		for i:=0;i < b.N; i++{
			c <- struct{}{}
		}
	}

	receiver := func(){
		defer wg.Done()
		<-begin
		for i:=0;i < b.N;i++{
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}