package _map

import (
	"testing"
)



// 并发读写 map 会出问题
// from:https://www.jianshu.com/p/10a998089486
// https://blog.golang.org/maps
func TestConcurrentMap(t *testing.T) {
	Map := make(map[int]int)

	for i := 0; i < 10; i ++ {
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}

}

func TestConcurrentMap2(t *testing.T) {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}

func readMap(Map map[int]int, key int) int {
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	Map[key] = value
}
