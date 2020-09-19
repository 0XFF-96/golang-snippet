package mockPkgErrors

import (
	"fmt"
	"strings"
)

// stack represents a stack of program counters.
type stack []uintptr


// st fmt.State 这个状态机函数应该好好看一下。
//
func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}

//func callers() *stack {
//	const depth = 32
//	var pcs [depth]uintptr
//	n := runtime.Callers(3, pcs[:])
//	var st stack = pcs[0:n]
//	return &st
//}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}