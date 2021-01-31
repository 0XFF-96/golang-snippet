package discovery

import (
	"context"
	"sync/atomic"

	"github.com/bilibili/discovery/conf"
	"github.com/bilibili/discovery/model"
	http "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// Discovery discovery.
type Discovery struct {
	c         *conf.Config

	// 自我保护机制
	// wiki:https://mp.weixin.qq.com/s/fr5m5x1wv8zQ4W2_po_9_g
	protected bool
	client    *http.Client
	registry  *Registry
	// 1. 为什么需要 atomic value ?
	// 2. 使用 atomic value 的好处是什么？
	nodes     atomic.Value
}

// New get a discovery.
func New(c *conf.Config) (d *Discovery, cancel context.CancelFunc) {
	return nil, nil
}
func (d *Discovery) exitProtect() {
}

// Register a new instance.
func (d *Discovery) Register(
	c context.Context,
	ins *model.Instance,
	latestTimestamp int64,
	replication bool,
	fromzone bool,
	) {
}

// Renew marks the given instance of the given app name as renewed,
// and also marks whether it originated from replication.
func (d *Discovery) Renew(
	c context.Context,
	arg *model.ArgRenew,
	) (i *model.Instance, err error) {
	return
}

// Cancel cancels the registration of an instance.
// 个人感觉这个名字叫 Revoke 会好一些。  - 取消
func (d *Discovery) Cancel(
	c context.Context,
	arg *model.ArgCancel,
	) (err error) {
	return
}

// FetchAll fetch all instances of all the department.
func (d *Discovery) FetchAll(c context.Context) (im map[string][]*model.Instance) {
	return nil
}

// Fetch fetch all instances by appid.
func (d *Discovery) Fetch(
	c context.Context,
	arg *model.ArgFetch,
	) (info *model.InstanceInfo, err error) {
	return nil, nil
}

// Fetchs fetch multi app by appids.
func (d *Discovery) Fetchs(
	c context.Context,
	arg *model.ArgFetchs,
	) (is map[string]*model.InstanceInfo, err error) {
	return
}
// Polls hangs request and then write instances when that has changes,
// or return NotModified.
func (d *Discovery) Polls(
	c context.Context,
	arg *model.ArgPolls,
	) (
	ch chan map[string]*model.InstanceInfo,
	new bool,
	miss []string,
	err error,
	) {
	return
}
// DelConns delete conn of host in appid
func (d *Discovery) DelConns(arg *model.ArgPolls) {
}
// Nodes get all nodes of discovery.
func (d *Discovery) Nodes(c context.Context) (nsi []*model.Node) {
	return nil
}

// Set set metadata,color,status of instance.
func (d *Discovery) Set(c context.Context, arg *model.ArgSet) (err error) {
	return nil
}
