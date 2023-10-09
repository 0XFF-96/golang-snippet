package main

import (
	"testing"
)

func TestMainFunc(t *testing.T) {
	//eb := &EventBus{
	//	subscribes: map[string]DataChannelSlice{},
	//	// rm:         sync.RWMutex{},
	//}
	//
	//
	//ch1 := make(chan DataEvent)
	//ch2 := make(chan DataEvent)
	//ch3 := make(chan DataEvent)
	//
	//eb.Subscribe("topic1", ch1)
	//eb.Subscribe("topic2", ch2)
	//eb.Subscribe("topic2", ch3)
	//
	//go publishTo("topic1", "Hi topic 1")
	//go publishTo("topic2", "Welcome to topic 2")
	//
	//for {
	//	select {
	//	case d := <-ch1:
	//		go printDataEvent("ch1", d)
	//	case d := <-ch2:
	//		go printDataEvent("ch2", d)
	//	case d := <-ch3:
	//		go printDataEvent("ch3", d)
	//	}
	//}
}
