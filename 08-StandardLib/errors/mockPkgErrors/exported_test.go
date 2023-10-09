package mockPkgErrors

import (
	"io"
	"testing"
)

// 编译的时候，某些地方出错了
// 是不是因为没有 main package ?
// 没有把这个目录下所有的代码都编译， 导致 undefined: Wrap !
//# command-line-arguments [command-line-arguments.test]
//./exported_test.go:15:4: undefined: Wrap
//./exported_test.go:19:10: undefined: Wrap

func TestWrap(t *testing.T) {
	tests := []struct {
		err     error
		message string
		want    string
	}{
		{io.EOF, "read error", "read error: EOF"},
		{Wrap(io.EOF, "read error"), "client error", "client error: read error: EOF"},
	}

	for _, tt := range tests {
		got := Wrap(tt.err, tt.message).Error()
		if got != tt.want {
			t.Errorf("Wrap(%v, %q): got: %v, want %v", tt.err, tt.message, got, tt.want)
		}
	}
}

// 探究一下 Cause 内部的原理
func TestCause2(t *testing.T) {
	got := Cause(New("error"))
	t.Log(got.Error())

	got = Cause(WithStack(New("error")))
	t.Log(got.Error())
}