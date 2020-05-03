package buildRedisJobQueue

import (
	"gopkg.in/redis.v3"
)

// Queue is the central element of this library.
// Packages can be put into or get from the queue.
// To read from a queue you need a consumer.
type Queue struct {
	redisClient    *redis.Client
	Name           string
	rateStatsCache map[int64]map[string]int64
	rateStatsChan  chan (*dataPoint)
	lastStatsWrite int64
}

// 重点函数
// newQueue
// Put、Delete、AddConsumer
// startStatsWriter

// 一个度量的表示
type dataPoint struct {
	name  string
	value int64
	incr  bool
}

// CreateQueue return a queue that you can Put() or AddConsumer() to
// Works like SelectQueue for existing queues
// TODO:@klaus 这个函数参数还可以进一步优化一下
func CreateQueue(redisHost, redisPort, redisPassword string, redisDB int64, name string) *Queue {
	return nil
}

// SelectQueue returns a Queue if a queue with the name exists
func SelectQueue(redisHost, redisPort, redisPassword string, redisDB int64, name string) (queue *Queue, err error) {
	return nil, nil
}

func newQueue(redisHost, redisPort, redisPassword string, redisDB int64, name string) *Queue {
	return nil
}

// Delete clears all input and failed queues as well as all consumers
// will not proceed as long as consumers are running
func (queue *Queue) Delete() error {
	return nil
}

// Put writes the payload into the input queue
func (queue *Queue) Put(payload string) error {
	return nil
}

// RequeueFailed moves all failed packages back to the input queue
func (queue *Queue) RequeueFailed() error {
	return nil
}

// ResetInput deletes all packages from the input queue
func (queue *Queue) ResetInput() error {
	return nil
}

// packages == payload ?
// ResetFailed deletes all [packages] from the failed queue
func (queue *Queue) ResetFailed() error {
	return nil
}

// GetInputLength returns the number of packages in the input queue
func (queue *Queue) GetInputLength() int64 {
	return 0
}

// 多个 consumers 的操作
func (queue *Queue) getConsumers() (consumers []string, err error) {
	return []string{}, nil
}

// AddConsumer returns a conumser that can write from the queue
func (queue *Queue) AddConsumer(name string) (c *Consumer, err error) {
	return nil, nil
}

// 统计用途的函数
func (queue *Queue) incrRate(name string, value int64) {
}

func (queue *Queue) startStatsWriter() {
	return
}

func (queue *Queue) writeStatsCacheToRedis(now int64) {
}

func (queue *Queue) isActiveConsumer(name string) bool {
	return false
}
