package main

//noinspection GoUnresolvedReference
import (
	"time"

	"github.com/chromedp/chromedp"
)

func text(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		// 访问页面
		chromedp.Navigate(`https://movie.douban.com/tag/`),
		// 等待列表渲染
		chromedp.Sleep(5 * time.Second),
		// 获取获取服务列表HTML
		chromedp.OuterHTML("#content", res, chromedp.ByID),
	}
}

//func main() {
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// 创建 chrome 实例
//	cdp, err := chromedp.New(ctx, chromedp.WithLog(log.Printf))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var res string
//
//	//执行调用
//	err = cdp.Run(ctx, text(&res))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 调用 Shutdown
//	err = cdp.Shutdown(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 等待 chrome 结束
//	err = cdp.Wait()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 正则匹配所要的内容
//	pattern := `class="tag">(.*?)</span>`
//	rp2 := regexp.MustCompile(pattern)
//	data := rp2.FindAllStringSubmatch(res, -1)
//
//	// 创建一个 txt 文件，写入获取的内容
//	f, err := os.Create("fenlei.txt")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	// 关闭 f
//	defer f.Close()
//
//	// 遍历切片，获取需要的内容，并写入 txt 文件
//	for i := 0; i < len(data); i++ {
//		fmt.Println(data[i][1])
//		f.WriteString(data[i][1] + "\n")
//	}
//
//}
