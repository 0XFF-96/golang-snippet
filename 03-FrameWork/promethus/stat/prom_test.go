package prom

import (
	"testing"
)


var errorsCount = BusinessErrCount

// PromError prometheus error count.
func PromError(name string, err error) {
	errorsCount.Incr(name)
}


func TestProm(t *testing.T) {
	err := Do()
	if err != nil {
		PromError("asdf", err)
	}
}

func Do() error {
	return nil
}
