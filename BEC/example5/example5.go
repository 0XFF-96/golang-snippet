// example5.go
package main

func fd(is []int, bs []byte) {
	if len(is) >= 256 {
		for _, n := range bs {
			_ = is[n] // 第7行： 需要边界检查
		}
	}
}

func fd2(is []int, bs []byte) {
	if len(is) >= 256 {
		is = is[:256] // 第14行： 一个暗示
		for _, n := range bs {
			_ = is[n] // 第16行： 边界检查消除了！
		}
	}
}

func fe(isa []int, isb []int) {
	if len(isa) > 0xFFF {
		for _, n := range isb {
			_ = isa[n & 0xFFF] // 第24行： 需要边界检查
		}
	}
}

func fe2(isa []int, isb []int) {
	if len(isa) > 0xFFF {
		isa = isa[:0xFFF+1] // 第31行： 一个暗示
		for _, n := range isb {
			_ = isa[n & 0xFFF] // 第33行： 边界检查消除了！
		}
	}
}

func main() {}
