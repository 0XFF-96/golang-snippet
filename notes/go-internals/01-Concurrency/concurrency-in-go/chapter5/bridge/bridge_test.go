package bridge

import (
	"fmt"
	"testing"
)


// 感觉不会有人想碰这些代码
//
func TestBridge(t *testing.T){
	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	)<-chan interface{} {
		valStream := make(chan interface{})
		go func(){
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream){
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	genVals := func() <-chan <-chan interface{}{
		chanStream := make(chan(<-chan interface{}))

		go func(){
			defer close(chanStream)
			for i:=0;i < 10;i++{
				stream := make(chan interface{}, 1)
				// 这三句话是什么意思呢？
				// 为什需要马上 close(stream) ?
				stream <-i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}
	for v := range bridge(nil, genVals()){
		fmt.Printf("%v ", v)
	}
}
