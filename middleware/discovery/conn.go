package discovery

import (
	"github.com/bilibili/discovery/model"
)

// conn the poll chan contains consumer.
type conn struct {
	ch         chan map[string]*model.InstanceInfo // TODO(felix): increase
	arg        *model.ArgPolls
	latestTime int64
	count      int
}

