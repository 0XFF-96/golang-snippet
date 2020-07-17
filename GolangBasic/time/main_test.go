package time

import (
	"fmt"
	"testing"
	"time"
)

var t, _  = time.Parse("2006-01-02", "2017-09-01")

func TestT(t *testing.T){
	a, _ := time.Parse("2006-01-02", "2017-09-01")
	b, _ := time.Parse("2006-01-02", "2020-06-28")
	d := a.Sub(b)

	t.Log(a.UnixNano())
	a = a.Add(8 * time.Hour)
	t.Log(a.UTC())
	fmt.Println(d.Hours() / 24)
	now := time.Now()
	t.Log(now.Sub(b).Minutes())
	t.Log(a.Sub(b).Minutes())

	loc, _ := time.LoadLocation("Asia/Shanghai")

	//loc, —— := time.LoadLocation("Beijing")
	//if err != nil {
	//	t.Log(err)
	//}
	time := time.Date(2020, 6, 29, 0, 0, 0, 0, loc)
	t.Log(time)


	t.Log(time.Before(now))
	t.Log(time.After(now))
}


