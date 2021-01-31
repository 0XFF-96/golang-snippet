package advance_concurrency

import (
	"time"
)

// converts Fetches to a stream
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item), // for Updates
	}
	go s.loop()
	return s
}

// sub implements the Subscription interface.
type sub struct {
	fetcher Fetcher   // fetches items
	updates chan Item // delivers items to the user
}

// loop fetches items using s.fetcher and sends them
// on s.updates.  loop exits when s.Close is called.

// 1. periodically call Fetch
// 2. send fetched items on the Updates channel
// 3. exit when Close is called, reporting any error

// BUG:
// 1. Bug 1: unsynchronized access to s.closed/s.err
// 2. time.Sleep may keep loop running
// 3. Bug 3: loop may block forever on s.updates
func (s *sub) loopV1() {
	for {
		if s.closed {
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err
			time.Sleep(10 * time.Second)
			continue
		}
		for _, item := range items {
			s.updates <- item
		}
		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now))
		}
	}
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	// TODO: make loop exit
	// TODO: find out about any error
	return nil
}


//func (s *naiveSub) Close() error {
//	s.closed = true
//	return s.err
//}

// 1. Close was called
// 2. it's time to call fetch
// 3. send an item on s.Update

// benefit
// 1. select lets loop avoid blocking indefinitely in any one state
// 2.

//func (s *sub) loop() {
//	... declare mutable state ...
//	for {
//		... set up channels for cases ...
//		select {
//		case <-c1:
//			... read/write state ...
//		case c2 <- x:
//			... read/write state ...
//		case y := <-c3:
//			... read/write state ...
//		}
//	}
//}

// case1: Close communicates with loop via s.closing.