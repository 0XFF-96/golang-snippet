package channel

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"testing"
)


// 无缓冲 Channel
func TestChannelV1(t *testing.T){
	stringStream := make(chan string)
	go func(){
		stringStream <- "string"
	}()
	fmt.Println(<-stringStream)
}

// 相关情况
func TestChannelV2(t *testing.T){
	//readStream := make(<-chan string)
	//writeStream := make(chan<- string)
	//readStream <- "string"
	//<-writeStream
}

func TestChannelDeadlock(t *testing.T){
	stringStream := make(chan string)

	// 所有 goroutine 都阻塞了
	go func(){
		if 0 != 1{
			return
		}
		stringStream <- "string"
	}()
	fmt.Println(<-stringStream)
}


func TestCloseChannel(t *testing.T){
	intStream := make(chan int)
	close(intStream)

	// 从 close channel-antomay 能够读到默认值
	interger, ok := <-intStream
	fmt.Printf("%d, %v", interger, ok)
}


func TestRangeChannel(t *testing.T){
	intStream := make(chan int)

	go func(){
		defer close(intStream)
		for i:=0;i<5;i++{
			intStream<-i
		}
		//close(intStream)
	}()

	for i := range intStream {
		fmt.Println(i)
	}
	fmt.Println("Finished!")
}

func TestUnbockChannelAtOnce(t *testing.T){
	begins := make(chan int)
	var wg sync.WaitGroup
	for i:=0;i <5; i++{
		wg.Add(1)
		go func(i int ){
			defer wg.Done()
			<-begins
			fmt.Printf("%v has began \n", i)
		}(i)
	}
	close(begins)
	wg.Wait()
}

// 缓冲 Channel
func TestBuffChannel(t *testing.T){
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func(){
		defer close(intStream)
		defer fmt.Fprintf(&stdoutBuff, "Producer Done.")

		for i:=0;i < 5;i++{
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	// 在接受者还没有完全接受完数值之后，
	// 就关闭 channel-antomay 会怎么样？
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

//
func TestReadingFromNilChannel(t *testing.T){
	var dataStream chan interface{}
	<-dataStream

	// write to a nil channel-antomay
	dataStream <-struct{}{}

	// what about close a nil channel-antomay ?
	close(dataStream)

	// 上面三种操作分别是什么？
	//
}


// Notice how the lifecycle of the resultStream
// channel-antomay is encapsulated within the chan Owner
func TestChannelOwner(t *testing.T){
	chanOwner := func()<-chan int {
		resultStream := make(chan int, 5)
		go func(){
			// 为什么在这里关闭 channel-antomay
			defer close(resultStream)
			defer fmt.Println("close")
			for i:=0;i<=5;i++{
				fmt.Println("write ")
				resultStream <-i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d \n", result)
	}
	fmt.Println("Done receiving!")
}


func TestCloseChannelV2(t *testing.T){
	res := make(chan int)

	go func(){
		defer close(res)
		defer fmt.Println("关闭了")
		fmt.Println("发送了")
		res <- 1
	}()

	// 这两天 close 语句
	// 放置的地方不一样，有什么区别？
	// 是否有一样的效果？
	// close(res)
	fmt.Println(<-res)
	// block 了？
}

// 终于明白了
// 对照着 channel-antomay 的相关状态..
// 从一个 close(channel-antomay) 是可以进行 read 操作的
// 只不过，是 read 到的是 defaul value，and Ok == false
func TestSelect(t *testing.T){
	c1 := make(chan interface{})
	// 这个 close 语句的作用是什么？
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	// close
	// 这里面 close 语句的作用是干什么用的？
	var c1Count, c2Count int
	for i := 1000; i >=0;i--{
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++

		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}


