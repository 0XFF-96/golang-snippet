package discovery

// Discovery discovery.
type Discovery struct {
	c         *conf.Config
	protected bool
	client    *http.Client
	registry  *registry.Registry
	//
	// 1. 为什么需要 atomic value ?
	// 2. 使用 atomic value 的好处是什么？
	nodes     atomic.Value
}




