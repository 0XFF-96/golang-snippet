package the_go_programming_language

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestLisenter(t *testing.T){
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn) // handle one connection at a time.
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func TestNet(t *testing.T){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil{
		log.Fatal(err)
	}
}

// P227
func TestNetCat3(t *testing.T){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout,conn) // we ignore err in this block.
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdout)
	conn.Close()
	<-done // wait for background goroutine to finish..
}

// pipeline1
func TestPipeLine1(t *testing.T){
	// naturals
	// Warning this program will not stop !! be carefule!
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func(){
		for x := 0; ; x++{
			naturals <- x
		}
	}()

	// Squarer
	go func(){
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (int main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

func TestAlternitiveSelect(t *testing.T){
	// 如何不加 1，这个参数会阻塞。因为 unbuffered chan 的读写必须同时发生
	// select 什么时候会跳出到外面...
	// 为什么改变 case 到顺序，不会改变程序到输出
	// default 什么时候会输出,保证 select 不会被阻塞
	// 什么时候会跳出循环？select ready
	// buffed chan ? unbuffered chan ?
	// unbuffered chan
	//同一时刻，同时有 读、写两端把持 channel。
	//如果只有读端，没有写端，那么 “读端”阻塞。
	//如果只有写端，没有读端，那么 “写端”阻塞
	// buffered chan

	ctx := context.Background()
	defer ctx.Done()
	t.Log("begin")
	ch := make(chan int,1)
	t.Log("log ch")
	for i:=0;i<10;i++{
		t.Log("for")
		// 在这里执行了多少步？
		select{
		case ch<-i:
			t.Logf("send num: %d",i)
		case x := <-ch:
			t.Log(x)
		//default:
		//	t.Log("d")
		}
	}
}
