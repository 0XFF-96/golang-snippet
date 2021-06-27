package proxy


// 如何利用 consul 访问服务？

// 步骤一
//active, err := xclient.NewActiveService(config.Consul.Address)
//if err != nil {
//return nil, err
//}

// 步骤二
//func NewActiveService(sd string) (*ActiveService, error) {
//	proxy, err := rpc.NewProxy("active", sd)
//	if err != nil {
//		return nil, err
//	}
//	client := active.NewActiveClient(proxy.Conn)
//	return &ActiveService{proxy: proxy, client: client}, nil
//}