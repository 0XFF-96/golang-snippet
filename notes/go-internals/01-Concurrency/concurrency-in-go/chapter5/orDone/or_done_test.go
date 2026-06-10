package orDone

import (
	"io"
	"testing"
)

// try readability first
// and avoid premature optimization
// or-done channel-antomay
// select default select
// 需要处理有错误、done、
// goroutine 被 canceled 的情况
func TestOrDone(t *testing.T) {
	// allow
	// for val := range orDone(done, myChan) {
	// Do something with val
	// }
}

var orDone = func(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	// for-select-for 的相关模式
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}


func performWrite(b *testing.B, writer io.Writer){
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, repeat(done, byte(0)), b.N){
		writer.Write([]byte{bt.(byte)})
	}
}
func tmpFileOrFatal()*os.File{
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("error:%v", err)
	}
	return file
}

func BenchmarkUnbuffedWrite(b *testing.B){
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B){
	bufferredFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferredFile))
}