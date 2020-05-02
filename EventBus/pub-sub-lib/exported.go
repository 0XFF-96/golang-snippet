package pub_sub_lib

// 有三种功能的带出方法:
// Sub:
// Pub:
// Close:

// New creates a new PubSub and starts
// a goroutine for handling operations
// The capcacity of the channels created by Sub
// and SubOnce will be as specified.?

func New(capacity int) *PubSub {
	return nil
}

// 返回一个 chan
// 当消息被投递到指定的 topics 时，
// 会通过 chan 返回

// Sub returns a channel on which messages published on any of
// the specified topics can be received.
func (ps *PubSub) Sub(topics ...string) chan interface{} {
	return nil
}

// SubOnce is similar to Sub, but only the first message published,?
// after subscription, on any of the specified topics can be received.
func (ps *PubSub) SubOnce(topics ...string) chan interface{} {
	return nil
}

// SubOnceEach returns a channel on which callers receive, at most, one message
// for each topic.
// 最多只有一条消息？
func (ps *PubSub) SubOnceEach(topics ...string) chan interface{} {
	return nil
}

// 通用的 sub 实现
func (ps *PubSub) sub(op operation, topics ...string) chan interface{} {
	return nil
}

// AddSub adds subscriptions to an existing channel.
func (ps *PubSub) AddSub(ch chan interface{}, topics ...string) {
}

// AddSubOnceEach adds subscriptions to an
// existing channel with SubOnceEach
func (ps *PubSub) AddSubOnceEach(ch chan interface{}, topics ...string) {}

func (ps *PubSub) Unsub(ch chan interface{}, topic ...string) {
}

// Pub Publishes the given message to all subscribers of the
// the specified topics.
func (ps *PubSub) Pub(msg interface{}, topics ...string) {
}

// TryPub publishes the given message to all subscribers of
// the specified topics if the topic has buffer space.
func (ps *PubSub) TryPub(msg interface{}, topics ...string) {
}

// closed 和 shutdown 两个有什么
// 不一样的含义？
func (ps *PubSub) Close(topics ...string) {
}

// Shutdown closes all subscribed channels and terminates the goroutine.
func (ps *PubSub) Shutdown() {
}
