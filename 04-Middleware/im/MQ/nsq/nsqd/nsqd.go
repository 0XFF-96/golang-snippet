package nsqd

import (
	"crypto/tls"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// 一共有这么多个参数
// 一个个来看，它们分别的作用是什么？
type NSQD struct {
	// 64bit atomic vars need to be first for proper alignment on 32bit platforms
	clientIDSequence int64

	sync.RWMutex

	opts atomic.Value

	// dl        *dirlock.DirLock
	isLoading int32
	errValue  atomic.Value
	startTime time.Time

	// topicMap map[string]*Topic

	clientLock sync.RWMutex
	// clients    map[int64]Client

	lookupPeers atomic.Value

	// tcpServer     *tcpServer
	tcpListener   net.Listener
	httpListener  net.Listener
	httpsListener net.Listener
	tlsConfig     *tls.Config

	poolSize int

	notifyChan           chan interface{}
	optsNotificationChan chan struct{}
	exitChan             chan int
	// waitGroup            util.WaitGroupWrapper
	// ci *clusterinfo.ClusterInfo
}
