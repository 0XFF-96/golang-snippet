package timewheel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// # command-line-arguments [command-line-arguments.test]
// 没有把 timewheel 目录下的几个文件，编译到一起。
func TestNewTimingWheel(t *testing.T) {

	_, err := NewTimingWheel(0, 10, func(key, value interface{}) {})
	assert.NotNil(t, err)
}

func TestA(t *testing.T){
	t.Log("A")
}
