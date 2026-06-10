package buildRedisJobQueue

import (
	"testing"

	. "github.com/matttproud/gocheck"
	"gopkg.in/redis.v3"
)

func Test(t *testing.T) {
	// 这种集成测试的方式，
	// 确实很不错
	TestingT(t)
}

type TestSuite struct {
	queue       *Queue
	consumer    *Consumer
	redisClient *redis.Client
}

var (
	redisHost     = "localhost"
	redisPort     = "6379"
	redisPassword = ""
	redisDB       = int64(9)

	// 从 go check 里面来的
	_ = Suite(&TestSuite{})
)
