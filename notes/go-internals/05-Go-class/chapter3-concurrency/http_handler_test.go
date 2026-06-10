package chapter3_concurrency

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
	"source-code-reading/go/src/net/http"
)

// 1. 阻塞方式奇怪
func TestHttp(t *testing.T){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil);err!= nil {
			log.Error()
		}
	}()

	select {}
}


// 2. 委派工作

func TestDelegation(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi")
	})
}
// 3. 一般抽象的写法
func serveAPP() {
}

func serveDebug() {
}

// 1. serveDebug 退出，代码仍然感知不到
// 2. serveApp 返回，将导致 main.main 退出。 只能
// 考 supervisor 的机制管理和启动
func main() {
	go serveDebug()
	// 阻塞
	serveAPP()
}