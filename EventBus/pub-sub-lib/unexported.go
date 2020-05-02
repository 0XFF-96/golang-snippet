package pub_sub_lib

// 一共有 5 个未导出方法
//
func (ps *PubSub) start() {}

func (reg *registry) add(topic string, ch chan interface{}, st subType) {}

func (reg *registry) send(topic string, msg interface{}) {}

func (reg *registry) sendNoWait(topic string, msg interface{}) {}

func (reg *registry) removeTopic(topic string) {}

func (reg *registry) removeChannel(ch chan interface{}) {}
