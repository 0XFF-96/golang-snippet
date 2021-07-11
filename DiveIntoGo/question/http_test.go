package question

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"testing"
)

// 为什么 body 需要 close
// https://segmentfault.com/a/1190000020086816?utm_source=sf-similar-article
// https://mp.weixin.qq.com/s/T6XXaFFyyOJioD6dqDJpFg
// https://www.cnblogs.com/lovezbs/p/13197587.html
// 参考 go/1.12.7/libexec/src/net/http/transport.go: 1758
// func (t *Transport) tryPutIdleConn(pconn *persistConn) error=
// func (es *bodyEOFSignal) Read(p []byte) (n int, err error) {
// https://segmentfault.com/a/1190000038267259?utm_source=sf-similar-article
func TestHttpClose(t *testing.T) {
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		_, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}

