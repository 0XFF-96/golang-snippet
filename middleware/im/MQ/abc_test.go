package MQ

import (
	"testing"
	"time"

	"github.com/sundayfun/daycam-server/util/grpcx"
)

func TestTB(t *testing.T){
	t.Log(grpcx.Timestamp(time.Now().UnixNano()))
	t.Log(grpcx.Timestamp(time.Now().AddDate(0, 0, -10).UnixNano()))
	t.Log(time.Now())
}
