package errors

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	t.Log(switcher(4, 5, 6))
	t.Log(switcher(0, 5, 6))
	t.Log(switcher(5, 5, 6))
	t.Log(switcher(2, 5, 6))
}

func switcher(a, b, c int) int {
	var answer int
	switch a {
	case 5:
	case 0:
	case 2:
		fmt.Println(2)
	case 7:
	case 4:
	default:
		answer = b
	}
	return answer
}