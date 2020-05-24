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
	ps := &PubSub{make(chan cmd), capacity}
	go ps.start()
	return ps
}

// 返回一个 chan
// 当消息被投递到指定的 topics 时，
// 会通过 chan 返回

// Sub returns a channel on which messages published on any of
// the specified topics can be received.
func (ps *PubSub) Sub(topics ...string) chan interface{} {
	return ps.sub(sub, topics...)
}

// SubOnce is similar to Sub, but only the first message published,?
// after subscription, on any of the specified topics can be received.
func (ps *PubSub) SubOnce(topics ...string) chan interface{} {
	return ps.sub(subOnce, topics...)
}

// SubOnceEach returns a channel on which callers receive, at most, one message
// for each topic.
// 最多只有一条消息？
// 搞不懂，
// 需要看看它是，如果处理 subOnce 的函数的
func (ps *PubSub) SubOnceEach(topics ...string) chan interface{} {
	return ps.sub(subOnce, topics...)
}

// 通用的 sub 实现
func (ps *PubSub) sub(op operation, topics ...string) chan interface{} {
	ch := make(chan interface{}, ps.capacity)
	ps.cmdChan <- cmd{op: op, topics: topics, ch: ch}
	// 把控制权给到外面
	// 这个 ch 在什么情况下，可以被关闭？
	// close ch ?
	//
	return ch
}

// AddSub adds subscriptions to an existing channel.
func (ps *PubSub) AddSub(ch chan interface{}, topics ...string) {
	ps.cmdChan <- cmd{op: sub, topics: topics, ch: ch}
}

// AddSubOnceEach adds subscriptions to an
// existing channel with SubOnceEach
func (ps *PubSub) AddSubOnceEach(ch chan interface{}, topics ...string) {
	ps.sub(subOnceEach, topics...)
}

func (ps *PubSub) Unsub(ch chan interface{}, topics ...string) {
	if len(topics) == 0 {
		ps.cmdChan <- cmd{op: unsubAll, ch: ch}
		return
	}

	ps.cmdChan <- cmd{op: unsub, topics: topics, ch: ch}
}

// Pub Publishes the given message to all subscribers of the
// the specified topics.
func (ps *PubSub) Pub(msg interface{}, topics ...string) {
	ps.cmdChan <- cmd{op: pub, topics: topics, msg: msg}
}

// TryPub publishes the given message to all subscribers of
// the specified topics if the topic has buffer space.
func (ps *PubSub) TryPub(msg interface{}, topics ...string) {
	ps.cmdChan <- cmd{op: tryPub, topics: topics, msg: msg}
}

// closed 和 shutdown 两个有什么
// 不一样的含义？
// 如果 close 之后
// 还继续用 ch 怎么办？
// 有办法处理这种情况吗？
func (ps *PubSub) Close(topics ...string) {
	ps.cmdChan <- cmd{op: closeTopic, topics: topics}
}

// Shutdown closes all subscribed channels and terminates the goroutine.
func (ps *PubSub) Shutdown() {
	ps.cmdChan <- cmd{op: shutdown}
}
