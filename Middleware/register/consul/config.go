package consul

const (
	envAddress    = "CONSUL_ADDRESS"
	envID         = "CONSUL_ID"
	envName       = "CONSUL_NAME"
	envTags       = "CONSUL_TAGS"
	envTTL        = "CONSUL_TTL"
	envToken      = "CONSUL_TOKEN"
	envDatacenter = "CONSUL_DATACENTER"
)

// FromEnv get config from env
func FromEnv() (*RegistryConfig, error) {
	return nil, nil
}

