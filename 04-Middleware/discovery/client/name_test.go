package client

// This Example register a server provider into discovery.
func ExampleDiscovery_Register() {
	// Unordered output4
}



type consumer struct {
	conf  *Config
	appID string
	dis   Resolver
	ins   []*Instance
}

func (c *consumer) getInstances(ch <-chan struct{}) {
	for { // NOTE: 通过watch返回的event chan =>
		if _, ok := <-ch; !ok {
			return
		}
		// NOTE: <= 实时fetch最新的instance实例
		ins, ok := c.dis.Fetch()
		if !ok {
			continue
		}
		// get local zone instances, otherwise get all zone instances.
		if in, ok := ins.Instances[c.conf.Zone]; ok {
			c.ins = in
		} else {
			for _, in := range ins.Instances {
				c.ins = append(c.ins, in...)
			}
		}
	}
}

func (c *consumer) getInstance() (ins *Instance) {
	// get instance by loadbalance
	// you can use any loadbalance algorithm what you want.
	return
}