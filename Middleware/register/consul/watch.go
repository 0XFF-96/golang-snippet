package consul

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/grpc/grpc-go/naming"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// 只暴露了 Close 和 Next 方法
type Watcher struct {
	sync.RWMutex
	serviceName string
	wp          *watch.Plan
	updates     chan []*naming.Update
	addrs       []*naming.Update
}

func newWatcher(serviceName string, address string) naming.Watcher {
	wp, err := watch.Parse(map[string]interface{}{
		"type":    "service",
		"service": serviceName,
	})

	if err != nil {
		return nil
	}

	w := &Watcher{
		serviceName: serviceName,
		wp:          wp,
		updates:     make(chan []*naming.Update),
	}
	wp.Handler = w.handle
	go wp.Run(address)
	return w
}

// Close close watcher
func (w *Watcher) Close() {
	w.wp.Stop()
	close(w.updates)
}

// Next next
func (w *Watcher) Next() ([]*naming.Update, error) {
	select {
	case updates := <-w.updates:
		return updates, nil
	}
}

func (w *Watcher) handle(idx uint64, data interface{}) {
	entries, ok := data.([]*api.ServiceEntry)
	if !ok {
		return
	}

	addrs := []*naming.Update{}
	for _, e := range entries {
		for _, check := range e.Checks {
			if check.ServiceID == e.Service.ID {
				if check.Status == api.HealthPassing {
					addr := fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port)
					metadata := map[string]string{}
					if len(e.Service.Tags) > 0 {
						err := json.Unmarshal([]byte(e.Service.Tags[0]), &metadata)
						if err != nil {
							fmt.Println("Parse node data error:", err)
						}
					}
					addrs = append(addrs, &naming.Update{Addr: addr, Metadata: &metadata})
				}
				break
			}
		}
	}

	updates := []*naming.Update{}

	// 这段添加和删除的代码，可以写得更加清晰一下
	// 尤其是，这两个逻辑是相同的，感觉可以用一个函数实现

	//add new address
	for _, newAddr := range addrs {
		found := false
		for _, oldAddr := range w.addrs {
			if newAddr.Addr == oldAddr.Addr {
				found = true
				break
			}
		}
		if !found {
			newAddr.Op = naming.Add
			updates = append(updates, newAddr)
		}
	}

	//delete old address
	for _, oldAddr := range w.addrs {
		found := false
		for _, addr := range addrs {
			if addr.Addr == oldAddr.Addr {
				found = true
				break
			}
		}
		if !found {
			oldAddr.Op = naming.Delete
			updates = append(updates, oldAddr)
		}
	}
	w.addrs = addrs

	w.updates <- updates
}
