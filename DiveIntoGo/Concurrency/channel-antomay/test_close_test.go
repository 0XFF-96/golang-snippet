package channel_antomay

import (
	"fmt"
	"testing"
)

// close stream
// 如何避免被重复 close ?
//func (s *sendLockServerStream) Close() {
//	if _, isClose := <-s.DoneCh; isClose {
//		return
//	}
//	close(s.DoneCh)
//}


// https://www.youtube.com/watch?v=ISfIT3B4y6E&ab_channel=PeriodicVideos
func TestClosed(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// iteration terminates after receving 3 values
	for elem := range c {
		fmt.Println(elem)
	}
}