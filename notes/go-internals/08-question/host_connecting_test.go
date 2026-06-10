package question

import (
	"fmt"
	"io"
	"net"
	"os"
	"testing"
)

// from https://mp.weixin.qq.com/s?__biz=MzAwMDAxNjU4Mg==&mid=2247485389&idx=1&sn=883fd935bc9a026dce30cbc1be532258&chksm=9aee2e64ad99a7723c06746e205b69205aabfda5ce8d8074192dd185e704e70a674b69f7eda4&scene=132#wechat_redirect
// 1. 连接一个 IP 不存在的主机时，握手过程是怎样的？
// 2. 图解网络硬件
// https://mp.weixin.qq.com/s?__biz=Mzg5NDY2MDk4Mw==&mid=2247486378&idx=1&sn=64058fa5061f6ab3e7e1d9626357cc96&source=41#wechat_redirect
//- 连接一个 IP 不存在的主机时，握手过程是怎样的？
//	1、分情况 局域网内 和 局域网外的
//  2、
//- 连接一个 IP 地址存在但端口号不存在的主机时，握手过程又是怎样的呢？
//- 正常情况的握手过程是怎么样的?


// 尝试根据这种情况，抓包就可以了。
// 1. 连接一个局域网内不存在的 ip
// 2. 没有三次握手的包，只有一些 ARP 包
// 3. arp -a 查看本地的 mac 地址
// 4. 如果目的IP跟本机IP不在同一个局域网下，那么会去获取默认网关的MAC地址

// 5. 局域网外的， 发出 TCP 第一次握手的 SYN包 。但是，会不断重试。

// 连IP 地址存在但端口号不存在的主机的握手过程
// TCP协议在识别到这个端口号对应的进程根本不存在时，
// 就会把数据丢弃，响应一个RST消息给发送端
func TestInlocal(t *testing.T) {
	client, err := net.Dial("tcp", "192.168.31.7:8081")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer client.Close()
	go func() {
		input := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(input)
			if err != nil {
				fmt.Println("input err:", err)
				continue
			}
			client.Write([]byte(input[:n]))
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := client.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read err:", err)
			continue
		}
		fmt.Println(string(buf[:n]))
	}
}


// 小结
//连一个 IP 不存在的主机时
//如果IP在局域网内，会发送N次ARP请求获得目的主机的MAC地址，同时不能发出TCP握手消息。
//如果IP在局域网外，会将消息通过路由器发出，但因为最终找不到目的地，触发TCP重试流程。
//连IP 地址存在但端口号不存在的主机时
//不管目的IP是回环地址还是局域网内外的IP地址，目的主机的传输层都会在收到握手消息后，发现端口不正确，发出RST消息断开连接。
//当然如果目的机器设置了防火墙策略，限制他人将消息发到不对外暴露的端口，那么这种情况，发送端就会不断重试第一次握手。
//最后留个问题，连一个 不存在的局域网外IP的主机时，我们可以看到TCP的重发规律是：开始时，每隔1s重发五次 TCP SYN消息，接着2s,4s,8s,16s,32s都重发一次；
//对比连一个 不存在的局域网内IP的主机时，却是每隔1s重发了4次ARP请求，接着过了32s后才再发出一次ARP请求。已知ARP请求是没有重传机制的，它的重试就是TCP重试触发的，但两者规律不一致，是为什么
