package heartbeat

import (
	"fmt"
	"testing"
	"time"
)

// What we discussed in the section
// 1、Systems saturation
// 2、Stale data
// 3、Attempting to prevent deadlocks
// 4、

// why a concurrent process might be canceled:
// 1、Timeouts
// 2、User intervention
// 3、Parent cancellation
// 4、Replicated requests : when the first come back,
// we may want to cancel the rest of the process.

// 是否可抢占
// 是否能够应对 upstream timeout 的情况

func TestDoWord(t *testing.T) {
	//doWord := func(
	//	done <-chan interface{},
	//	pulseInterval time.Duration,
	//) (<-chan interface{}, <-chan time.Time) {
	//	heartbeat := make(chan interface{})
	//	results := make(chan time.Time)
	//
	//	go func() {
	//		defer close(heartbeat)
	//		defer close(results)
	//
	//		pulse := time.Tick(pulseInterval)
	//		workGen := time.Tick(2 * pulseInterval)
	//
	//		sendPulse := func() {
	//			select {
	//			case heartbeat <- struct{}{}:
	//			default:
	//			}
	//		}
	//		sendResult := func(r time.Time) {
	//			for {
	//				select {
	//				case <-done:
	//					return
	//				case <-pulse:
	//					sendPulse()
	//				case results <- r:
	//					return
	//				}
	//			}
	//		}
	//	}
	//
	//}
}

func TestHeadBeat(t *testing.T) {
	doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		results := make(chan time.Time)

		go func() {
			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default:

				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}

	done := make(chan interface{})
	const timeout = 2  * time.Second
	time.AfterFunc(10 *time.Second, func(){close(done)})
	heartbeat, results  := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("result %v\n", r)

		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}
