package discovery

import (
	"context"
	"net/http"

	"github.com/bilibili/discovery/conf"
	"github.com/bilibili/discovery/model"
)

// Node represents a peer node to which information should be shared from this node.
//
// This struct handles replicating all update operations like 'Register,Renew,Cancel,
//  Expiration and Status Changes'
// to the <Discovery Server> node it represents.
type Node struct {
	c *conf.Config

	// url
	client       *http.Client
	pRegisterURL string
	registerURL  string
	cancelURL    string
	renewURL     string
	setURL       string

	addr      string
	status    model.NodeStatus
	zone      string
	otherZone bool
}


// newNode return a node.
func newNode(c *conf.Config, addr string) (n *Node) {
	return
}

// Register send the registration information of
// Instance receiving by this node to the peer node represented.
func (n *Node) Register(c context.Context, i *model.Instance) (err error) {
	return nil
}

// Cancel send the cancellation information of Instance
// receiving by this node to the peer node represented.
func (n *Node) Cancel(c context.Context, i *model.Instance) (err error) {
	return
}

// Renew send the heartbeat information of
// Instance receiving by this node to the peer node represented.
// If the instance does not exist the node,
// the instance registration information is sent again to the peer node.
func (n *Node) Renew(c context.Context, i *model.Instance) (err error) {
	return
}

// Set the infomation of instance by this node to the peer node represented
func (n *Node) Set(c context.Context, arg *model.ArgSet) (err error) {
	return nil
}
func (n *Node) call(
	c context.Context,
	action model.Action,
	i *model.Instance,
	uri string,
	data interface{},
	) (err error) {
	return
}

func (n *Node) setCall(
	c context.Context,
	arg *model.ArgSet,
	uri string,
	) (err error) {
	return
}

