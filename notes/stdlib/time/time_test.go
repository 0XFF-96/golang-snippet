package time

import (
	"testing"
	"time"
)

// https://golangtc.com/t/54f6c643421aa9089a00002f
// https://studygolang.com/articles/8042
func TestTodayTime(t *testing.T) {
	now := time.Now().UTC()

	// 每天 7 点执行一次
	// 判断是否 23 点
	// 16:00
	// 15:00 - (+7) 23:00 (这段时间）

	// 16:00 + 7 = 23
	t.Log(time.Now().UTC())
	// 当前时间的 23 点

	// 是否今天的 7 点，不是的话，跳过。

	// 七点只推最后一条消息
	// map[uint64]*Item
	// lock
	// 每天七点左右，一次过取走 map 里面的所有 item, 然后进行推送
	t.Log(now.Hour())
	t.Log(now.Day())
	t.Log(time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, time.UTC))
	t.Log(time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, time.UTC))
}
