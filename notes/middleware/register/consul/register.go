package consul

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

// Registry consul registry
type Registry struct {
	closeCh chan struct{}
	client  *api.Client
	cfg     *RegistryConfig
}

// NewRegistry new registry
func NewRegistry(cfg *RegistryConfig) (*Registry, error) {
	consulCfg := api.Config{
		Address:    cfg.Address,
		Token:      cfg.Token,
		Datacenter: cfg.Datacenter,
	}
	c, err := api.NewClient(&consulCfg)
	if err != nil {
		return nil, err
	}

	return &Registry{
		client:  c,
		cfg:     cfg,
		closeCh: make(chan struct{}),
	}, nil
}

func (r *Registry) registeAndCheck(agent *api.Agent, serviceID string) error {
	metadata, err := json.Marshal(r.cfg.Meta)
	if err != nil {
		return err
	}
	tags := make([]string, 0)
	tags = append(tags, string(metadata))

	reg := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    r.cfg.Name,
		Address: r.cfg.NodeAddress,
		Port:    r.cfg.NodePort,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			TTL:                            fmt.Sprintf("%ds", r.cfg.TTL),
			Status:                         api.HealthPassing,
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	err = agent.ServiceRegister(reg)
	if err != nil {
		return fmt.Errorf("ServiceRegister error: %s", err.Error())
	}

	check := api.AgentServiceCheck{
		TTL:    fmt.Sprintf("%ds", r.cfg.TTL),
		Status: api.HealthPassing,
	}
	err = agent.CheckRegister(
		&api.AgentCheckRegistration{
			ID:                "service:" + serviceID,
			Name:              r.cfg.Name,
			ServiceID:         serviceID,
			AgentServiceCheck: check,
		},
	)
	if err != nil {
		return fmt.Errorf("CheckRegister error: %s", err.Error())
	}

	return nil
}

// Register register
func (r *Registry) Register() error {
	serviceID := fmt.Sprintf("%s-%s-%d", r.cfg.Name, r.cfg.NodeAddress, r.cfg.NodePort)

	agent := r.client.Agent()
	err := r.registeAndCheck(agent, serviceID)
	if err != nil {
		return err
	}

	keepAliveTicker := time.NewTicker(time.Duration(r.cfg.TTL) * time.Second / 3)
	for {
		select {
		case <-r.closeCh:
			keepAliveTicker.Stop()
			agent.ServiceDeregister(serviceID)
			return nil
		case <-keepAliveTicker.C:
			err := agent.UpdateTTL("service:"+serviceID, "", api.HealthPassing)
			if err != nil {
				fmt.Printf("UpdateTTL error: %s\n", err.Error())
				r.registeAndCheck(agent, serviceID)
			}
		}
	}
}

// Deregister de register
func (r *Registry) Deregister() error {
	r.closeCh <- struct{}{}
	return nil
}
