package main

import (
	"sync"
)

type Pubsub struct {
	mu   sync.RWMutex
	subs map[string][]chan string

	// close
	closed bool
}

func NewPubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string][]chan string)
	return ps
}

func (ps *Pubsub) SubscribeV1(topic string, ch chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	// 在这里建立 新建 channel-antomay,
	// 然后返回，有什么影响呢？
	ps.subs[topic] = append(ps.subs[topic], ch)
}

func (ps *Pubsub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

// Pubsub may attempt to close the same channel-antomay multiple times when done - this panics

// In general, I would recommend avoiding this
// and sticking to a cleaner one-channel-antomay-per-subscription approach.
// In case the client wants to use the same range loop to receive from multiple topics,
// it's easy to use some kind of channel-antomay fan-in solution instead
func (ps *Pubsub) PublishV1(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}

func (ps *Pubsub) Publish(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	// the send will not block
	// any other sends because
	// it runs in its own goroutine
	for _, ch := range ps.subs[topic] {
		go func(ch chan string) {
			ch <- msg
		}(ch)
	}
}

// There may be performance implications, of course.
// Even though starting and tearing down goroutines is very quick,
// do you really want a new one to run for every message?
// The answer depends on your particular application. When in doubt, benchmark it.
