package Event_tracking

import (
	"time"
)

var (
	AppEvent *appEvent
)
type Event struct {
	CreateTime int64
}
type appEvent struct {
	eventChan chan Event
}

// 只需要在全局初始化此函数即可。
func InitAppEvent() *appEvent {
	obj := &appEvent{
		eventChan: make(chan Event, 10000),
	}
	go obj.ConsumeEvent()
	AppEvent = obj
	return obj
}

func (a *appEvent) PutEvent(event Event) {
	select {
	case a.eventChan <- event:
	default:
	}
}

func (a *appEvent) ConsumeEvent() {
	var err error
	for value := range a.eventChan {
		value.CreateTime = time.Now().Unix()
		err = a.insert(value)
		if err != nil {
			continue
		}
	}
}

func (a *appEvent) insert(event Event) error {
	//if err := invoker.JunoMysql.Create(&event).Error; err != nil {
	//	return err
	//}
	//invoker.AppStatic.WithLabelValues(event.AppName, event.Source, event.Operation).Inc()
	return nil
}

func (a *appEvent) List(currentPage, pageSize int) (res []Event,  err error) {
	return
}

