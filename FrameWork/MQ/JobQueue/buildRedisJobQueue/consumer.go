package buildRedisJobQueue

import (
	"context"

	"gopkg.in/redis.v3"
)

// Consumer are used for reading from queues
type Consumer struct {
	Name  string
	Queue *Queue

	// 这个 cancel 的处理有点奇怪
	// 为什么这样用？
	cancel         context.CancelFunc
	contextCleared <-chan struct{}
}

/*       **/
// Get returns a single package from the queue (blocking)
func (consumer *Consumer) Get() (*Package, error) {
	return nil, nil
}

// NoWaitGet returns a single package from the queue (returns nil, nil if no package in queue)
// BlockedGet
// UnblockGet
//  这样名明好像更好一下？
func (consumer *Consumer) NoWaitGet() (*Package, error) {
	return nil, nil
}

// MultiGet returns an array of packages from the queue
func (consumer *Consumer) MultiGet(length int) ([]*Package, error) {
	return nil, nil
}

// GetUnacked returns a single packages from the working queue of this consumer
func (consumer *Consumer) GetUnacked() (*Package, error) {
	return nil, nil
}

/*       **/
// HasUnacked returns true if the consumers has unacked packages
func (consumer *Consumer) HasUnacked() bool {
	return false
}

// GetUnackedLength returns the number of packages in the unacked queue
func (consumer *Consumer) GetUnackedLength() int64 {
	return 0
}

// GetFailed returns a single packages from the failed queue of this consumer
func (consumer *Consumer) GetFailed() (*Package, error) {
	return nil, nil
}

/*       **/

// ResetWorking deletes! all messages in the working queue of this consumer
// 这个操作我们没有！！
//
func (consumer *Consumer) ResetWorking() error {
	return nil
}

// RequeueWorking requeues all packages from working to input
func (consumer *Consumer) RequeueWorking() error {
	return nil
}

func (consumer *Consumer) ackPackage(p *Package) error {
	return nil
}

func (consumer *Consumer) requeuePackage(p *Package) error {
	return nil
}

func (consumer *Consumer) failPackage(p *Package) error {
	return nil
}

// 为什么这里进行 startHeartbeat ?
// 我们的函数不需要？
func (consumer *Consumer) startHeartbeat() {
}

func (consumer *Consumer) Quit() {

}

func (consumer *Consumer) parseRedisAnswer(answer *redis.StringCmd) (*Package, error) {
	return nil, nil
}

func (consumer *Consumer) unsafeGet() (*Package, error) {
	return nil, nil
}
