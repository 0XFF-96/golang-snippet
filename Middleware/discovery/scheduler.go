package discovery

import (
	"sync"

	"github.com/bilibili/discovery/model"
)

// Scheduler info.
type scheduler struct {
	schedulers map[string]*model.Scheduler
	mutex      sync.RWMutex
	r          *Registry
}


func newScheduler(r *Registry) *scheduler {
	return nil
}

// Load load scheduler info.
func (s *scheduler) Load() {
}

// Reload reload scheduler info.
func (s *scheduler) Reload() {
}

// Get get scheduler info.
func (s *scheduler) Get(appid, env string) *model.Scheduler {
}