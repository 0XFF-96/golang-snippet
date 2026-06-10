package main

func f5(s []int) {
	for i := range s {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f6(s []int) {
	for i := 0; i < len(s); i++ {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f7(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		_ = s[i]   // 在 Go SDK 1.11 之前，这行仍旧需要边界检查
		_ = s[i:len(s)]
	}
}

func f8(s []int, index int) {
	if index >= 0 && index < len(s) {
		_ = s[index]
		_ = s[index:len(s)]
	}
}

func f9(s []int) {
	if len(s) > 2 {
		_, _, _ = s[0], s[1], s[2]
	}
}

func main() {}



