package buildRedisJobQueue

import (
	"time"
)

// Package provides headers and handling functions around payloads
// ？Collection 是干什么用的？
// 为什么需要把  queue 和 consumer  放在这个结构里面？
// 传递数据，redis 的数据格式。
type Package struct {
	Payload   string
	CreatedAt time.Time

	// ?
	Queue      interface{} `json:"-"`
	Consumer   *Consumer   `json:"-"`
	Collection *[]*Package `json:"-"`
	// ?
	Acked bool `json:"-"`
	//TODO add Headers or smth. when needed
	//wellle suggested error headers for failed packages
}

func unmarshalPackage(input string, queue *Queue, consumer *Consumer) (*Package, error) {
	return nil, nil
}

func (pack *Package) getString() string {
	return ""
}

func (pack *Package) index() int {
	return 0
}

// MultiAck removes all packaes from the fetched array up to and including this package
func (pack *Package) MultiAck() (err error) {
	return nil
}

// 为什么是一个包？
// Ack removes the packages from the queue
func (pack *Package) Ack() error {
	return nil
}

// Requeue moves a package back to input
func (pack *Package) Requeue() error {
	return nil
}

// Fail moves a package to the failed queue
func (pack *Package) Fail() error {
	return nil
}

func (pack *Package) reject(requeue bool) error {
	return nil
}
