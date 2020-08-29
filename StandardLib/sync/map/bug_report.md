```
(~/src) % env GODEBUG=gctrace=1 ./leak
gc 1 @0.015s 4%: 0.024+3.3+0.018 ms clock, 0.096+0.37/3.1/0.077+0.074 ms cpu, 5->5->3 MB, 6 MB goal, 4 P
gc 2 @0.032s 5%: 0.009+4.8+0.031 ms clock, 0.038+0.45/4.3/5.4+0.12 ms cpu, 8->8->6 MB, 9 MB goal, 4 P
gc 3 @0.069s 6%: 0.014+11+0.017 ms clock, 0.058+0.31/10/7.3+0.070 ms cpu, 16->17->12 MB, 17 MB goal, 4 P
gc 4 @0.152s 5%: 0.012+34+0.018 ms clock, 0.050+4.6/16/16+0.074 ms cpu, 33->34->25 MB, 34 MB goal, 4 P
gc 5 @0.317s 5%: 0.013+48+0.018 ms clock, 0.055+3.8/37/63+0.075 ms cpu, 67->67->50 MB, 68 MB goal, 4 P
gc 6 @0.652s 5%: 0.014+92+0.052 ms clock, 0.058+0.66/91/157+0.20 ms cpu, 135->135->100 MB, 136 MB goal, 4 P
gc 7 @1.322s 6%: 0.014+203+0.020 ms clock, 0.057+42/192/357+0.080 ms cpu, 269->270->200 MB, 270 MB goal, 4 P
gc 8 @2.721s 7%: 0.014+479+0.019 ms clock, 0.056+120/478/955+0.079 ms cpu, 539->541->400 MB, 540 MB goal, 4 P
gc 9 @5.619s 10%: 0.016+1205+0.026 ms clock, 0.064+759/1202/2360+0.10 ms cpu, 1079->1081->799 MB, 1080 MB goal, 4 P
^C
(~/src) % cat leak.go 
package main

import "sync"

func main() {
	var sm sync.Map

	var value [16]byte

	for i := 0; i < 1<<26; i++ {
		sm.Store(i, value)
		sm.Delete(i - 1)
	}
}

```

// 这个 PR 记录📝这个 Bug 修复的过程。 
// https://github.com/golang/go/commit/19c8546cd9cd53da0f4604d20f47714554fed051