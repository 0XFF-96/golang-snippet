package mockPkgErrors

import (
	"runtime"
)

// fundamental is an error that has a message and a stack, but no caller.
type fundamental struct {
	msg string
	*stack
}


func New(message string) error {
	return &fundamental{
		msg:   message,
		stack: callers(),
	}
}

// 实现了这个接口，就能实现 error 的接口方法
func (f *fundamental) Error() string { return f.msg }

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr

	// 在这里直接转换成了 切片类型
	// Cannot use 'pcs' (type [32]uintptr) as type []uintptr
	// n := runtime.Callers(3, pcs)
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
