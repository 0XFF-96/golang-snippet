package lb


type baseSelector struct {
	addrs   []string
	addrMap map[string]*AddrInfo
}

