package consul

// RegistryConfig registry config
// 疑惑🤔
// 这里为什么没有 zone、region、hostname 之类的信息呢？
// 这里是 consul 的配置...!
type RegistryConfig struct {
	Address     string            `yaml:"address"`
	Name        string            `yaml:"name"`
	Tags        []string          `yaml:"tags"`
	NodePort    int               `yaml:"nodePort"`
	NodeAddress string            `yaml:"nodeAddress"`
	Meta        map[string]string `yaml:"meta"`
	ID          string            `yaml:"id"`

	TTL        int    `yaml:"ttl"`
	Token      string `yaml:"token"`
	Datacenter string `yaml:"datacenter"`
}

