package pub_sub_lib

// 一共有 5 个未导出方法
//
// 不能导出的方法
// TODO: @klaus
// 能不能消除导出方法？
// 核心处理函数
// 事件中心的处理方式
func (ps *PubSub) start() {
	reg := registry{
		topics:    make(map[string]map[chan interface{}]subType),
		revTopics: make(map[chan interface{}]map[string]bool),
	}

loop:
	for cmd := range ps.cmdChan {
		if cmd.topics == nil {
			switch cmd.op {
			case unsubAll:
				reg.removeChannel(cmd.ch)

			case shutdown:
				break loop
			}

			continue loop
		}

		for _, topic := range cmd.topics {
			switch cmd.op {
			case sub:
				reg.add(topic, cmd.ch, normal)

			case subOnce:
				reg.add(topic, cmd.ch, onceAny)

			case subOnceEach:
				reg.add(topic, cmd.ch, onceEach)

			case tryPub:
				reg.sendNoWait(topic, cmd.msg)

			case pub:
				reg.send(topic, cmd.msg)

			case unsub:
				reg.remove(topic, cmd.ch)

			case closeTopic:
				reg.removeTopic(topic)
			}
		}
	}

	for topic, chans := range reg.topics {
		for ch := range chans {
			reg.remove(topic, ch)
		}
	}
}

func (reg *registry) add(topic string, ch chan interface{}, st subType) {}

func (reg *registry) send(topic string, msg interface{}) {}

func (reg *registry) sendNoWait(topic string, msg interface{}) {}

func (reg *registry) removeTopic(topic string) {}

func (reg *registry) removeChannel(ch chan interface{}) {}

func (reg *registry) remove(topic string, ch chan interface{}) {
}
