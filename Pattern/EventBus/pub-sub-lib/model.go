package pub_sub_lib

type operation int

const (
	sub operation = iota
	subOnce
	subOnceEach
	pub
	tryPub
	unsub
	unsubAll
	closeTopic
	shutdown
)

// PubSub is a collection of topics.
// topics 管理的集合
type PubSub struct {
	cmdChan  chan cmd
	capacity int
}

type cmd struct {
	op     operation
	topics []string
	ch     chan interface{}

	// 为什么 msg 定义为一个 interface ?
	msg interface{}
}

// registry maintains the current subscription state. It's not
// safe to access a registry from multiple goroutines simultaneously.

// registry 是一个全局变量
// 它维护当前订阅的状态（更好的表达？）
// 它不是一个线程安全变量
type registry struct {
	topics map[string]map[chan interface{}]subType

	revTopics map[chan interface{}]map[string]bool
}

type subType int

const (
	onceAny subType = iota
	onceEach
	normal
)
