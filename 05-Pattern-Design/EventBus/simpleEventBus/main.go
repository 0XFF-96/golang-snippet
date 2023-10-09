package main

import (
	"fmt"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

type EventBus struct {
	// A topic is treated as a key inside the map
	subscribes map[string]DataChannelSlice
	rm         sync.RWMutex
}

// Subscribe
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	fmt.Printf("subscribe %s", topic)
	eb.rm.Lock()
	if prev, found := eb.subscribes[topic]; found {
		eb.subscribes[topic] = append(prev, ch)
	} else {
		eb.subscribes[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	fmt.Println("Publishing")
	eb.rm.RLock()
	if chans, found := eb.subscribes[topic]; found {
		fmt.Println("found")
		channels := append(DataChannelSlice{}, chans...)

		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

// used-for test
//var eb = &EventBus{
//	subscribes: map[string]DataChannelSlice{},
//	// 原 pos 少了这个...
//	rm:         sync.RWMutex{},
//}

func (eb *EventBus) publishTo(topic string, data string) {
	for {
		// 这样不会死循环吗？
		fmt.Println("publish")
		eb.Publish(topic, data)
		time.Sleep(1 * time.Second)
		// time.Sleep(time.Duration(rand.Intn(10)) *time.Microsecond)
	}
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s Topic: %s; DataEvent: %v \n",
		ch,
		data.Topic,
		data.Data,
	)
}

func main() {
	eb := &EventBus{
		subscribes: map[string]DataChannelSlice{},
		// rm:         sync.RWMutex{},
	}

	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go eb.publishTo("topic1", "Hi topic 1")
	go eb.publishTo("topic2", "Welcome to topic 2")

	fmt.Println("start handling")
	for {
		select {
		case d := <-ch1:
			fmt.Println("ch1")
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}
