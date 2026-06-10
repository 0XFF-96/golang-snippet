package discovery

import (
	"context"

	"github.com/bilibili/discovery/conf"
	"github.com/bilibili/discovery/model"
	"go-common/library/sync/errgroup"
)

// Nodes is helper to manage lifecycle of a collection of Nodes.
type Nodes struct {
	nodes    []*Node
	zones    map[string][]*Node
	selfAddr string
}


// NewNodes new nodes and return.
func NewNodes(c *conf.Config) *Nodes {
	return &Nodes{
	}
}

// Replicate replicate information to all nodes except for this node.
func (ns *Nodes) Replicate(
	c context.Context,
	action model.Action,
	i *model.Instance,
	otherZone bool,
	) (err error) {
	return
}

// ReplicateSet replicate set information to all nodes except for this node.
func (ns *Nodes) ReplicateSet(
	c context.Context,
	arg *model.ArgSet,
	otherZone bool,
	) (err error) {
	return
}
func (ns *Nodes) action(
	c context.Context,
	eg *errgroup.Group,
	action model.Action,
	n *Node,
	i *model.Instance,
	) {
}

// Nodes returns nodes of local zone.
func (ns *Nodes) Nodes() (nsi []*Node) {
	return
}

// AllNodes returns nodes contain other zone nodes.
func (ns *Nodes) AllNodes() (nsi []*Node) {
	return
}

// Myself returns whether or not myself.
func (ns *Nodes) Myself(addr string) bool {
}

// UP marks status of myself node up.
func (ns *Nodes) UP() {
}