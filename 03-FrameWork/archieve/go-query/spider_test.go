package go_query

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func TestCrawl(t *testing.T) {
	resp, err := http.Get("https://s.weibo.com/top/summary?Refer=top_hot&topnav=1&wvr=6")
	if err != nil {
		t.Log(err)
	}

	var ret []string
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		t.Log(err)
	}
	doc.Find("table tr td").Each(func(i int, selection *goquery.Selection) {
		ret = append(ret, selection.Find("a").Text())
	})
	fmt.Printf("%s", strings.Join(ret, "\n"))
}


func TestCrawlT(t *testing.T) {
	// request, err := http.NewRequest("GET", "https://book.douban.com/chart?subcat=I", nil)
	request, err := http.NewRequest("GET", "https://book.douban.com/chart?subcat=F", nil)
	if err != nil {
		t.Log(err)
	}
	request.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) `+
		`AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36`)

	crawlerClient := http.Client{Timeout: 15 * time.Second}
	resp, err := crawlerClient.Do(request)
	if err != nil {
		t.Log()
	}
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	doc.Find(".article").Each(func(i int, selection *goquery.Selection) {
		selection.Find("li").Each(func(i int, selection *goquery.Selection) {
			name := selection.Find("a").Text()
			name = strings.TrimSpace(strings.Split(name, " ")[0])
			// fmt.Println(strings.TrimSpace(strings.Split(name, " ")[0]))
			author := selection.Find(".subject-abstract").Text()
			author = strings.TrimSpace(strings.Split(author, "/")[0])
			// fmt.Println(strings.TrimSpace(strings.Split(author, "/")[0]))
			fmt.Printf("%s,%s\n",name, author)
		})
	})
}

func TestDoubanMovie(t *testing.T){
	request, err := http.NewRequest("GET", "https://movie.douban.com/chart", nil)
	if err != nil {
		t.Log(err)
	}
	request.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) `+
		`AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36`)

	crawlerClient := http.Client{Timeout: 15 * time.Second}
	resp, err := crawlerClient.Do(request)
	if err != nil {
		t.Log()
	}
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	doc.Find("table").Each(func(i int, selection *goquery.Selection) {
		// movie := selection.Find(".pl2").Text()
		movie := strings.Split(selection.Find(".pl2").Text(), "/")[0]
		authors := strings.Split(selection.Find(".pl").Text(),"/")
		var ret []string
		for _, a := range authors[1:7] {
			ret = append(ret, strings.TrimSpace(strings.Replace(a, "·", "", -1)))
		}
		fmt.Printf("%s,%s\n",strings.TrimSpace(movie), strings.Join(ret, ","))
	})
}

// 这个不行
func TestDoubanWeekly(t *testing.T){
	request, err := http.NewRequest("GET", "https://m.douban.com/subject_collection/tv_global_best_weekly", nil)
	if err != nil {
		t.Log(err)
	}
	request.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) `+
		`AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36`)

	crawlerClient := http.Client{Timeout: 15 * time.Second}
	resp, err := crawlerClient.Do(request)
	if err != nil {
		t.Log()
	}
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(doc.Text())
	doc.Find("#virtual-item-0").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func estABC(t *testing.T){
	// https://mojotv.cn/2018/12/26/chromedp-tutorial-for-golang
	var html string
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	err := chromedp.Run(
		ctx,
		chromedp.Navigate("https://movie.douban.com/tv/#!type=tv&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=240"),
		chromedp.WaitVisible(`#footer`, chromedp.ByID),
		chromedp.OuterHTML("#body", &html),
	)
	t.Log(err)
	t.Log(html)
}

func TestBB(t *testing.T){


}