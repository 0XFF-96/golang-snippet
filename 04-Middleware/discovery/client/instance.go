package client

// InstancesInfo instance info.
type InstancesInfo struct {
	Instances map[string][]*Instance `json:"instances"`
	LastTs    int64                  `json:"latest_timestamp"`
	Scheduler []Zone                 `json:"scheduler"`
}

// Instance represents a server the client connects to.
type Instance struct {
	// Region is region.
	Region string `json:"region"`
	// Zone is IDC.
	Zone string `json:"zone"`
	// Env prod/pre、uat/fat1
	Env string `json:"env"`
	// AppID is mapping servicetree appid.
	AppID string `json:"appid"`
	// Hostname is hostname from docker.
	Hostname string `json:"hostname"`
	// Addrs is the address of app instance
	// format: scheme://host
	Addrs []string `json:"addrs"`
	// Version is publishing version.
	Version string `json:"version"`
	// LastTs is instance latest updated timestamp
	LastTs int64 `json:"latest_timestamp"`
	// Metadata is the information associated with Addr, which may be used
	// to make load balancing decision.
	Metadata map[string]string `json:"metadata"`
}

// UseScheduler use scheduler info on instances.
// if instancesInfo contains scheduler info about zone,
// return releated zone's instances weighted by scheduler.
// if not,only zone instances be returned.
func (insInf *InstancesInfo) UseScheduler(
	zone string,
	) (inss []*Instance) {
	return
}
