package pustToSessionbug

import (
	"fmt"
	"testing"
	"time"
)

type item struct {
	tilte string
}

func TestCopy(t *testing.T) {
	n := &item{tilte: ""}
	//for _, os := range []string{"ios", "android"}{
	//	go pushToDiffentDevice(n, os)
	//}

	for _, os := range []string{"android", "ios"} {
		go pushToDiffentDevice(n, os)
	}

	// 解决办法？
	// 把 xxx 变为 ?
	// 还有没有更好的办法呢？
	time.Sleep(1 * time.Second)
}

func pushToDiffentDevice(n *item, os string) {
	// nCopy := n  // 这个解决版本不行，持有的仍然是 操作 n 的指针。

	switch os {
	// fmt.Println(n.tilte)
	case "ios":
		pushIOS(*n)
	case "android":
		pushAndroid(*n)
	default:

	}
}

func pushAndroid(n item) {
	n.tilte = "have title"
	fmt.Println("android", n.tilte)
}

func pushIOS(n item) {
	fmt.Println("ios:", n.tilte)
}
