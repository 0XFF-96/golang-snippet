package _map

import (
	"sync"
	"sync/atomic"
)

// DeepCopyMap is an implementation of mapInterface using a Mutex and
// atomic.Value.  It makes deep copies of the map on every write to avoid
// acquiring the Mutex in Load.
type DeepCopyMap struct {
	mu    sync.Mutex
	clean atomic.Value
}

func (m *DeepCopyMap) Load(key interface{}) (value interface{}, ok bool) {
	clean, _ := m.clean.Load().(map[interface{}]interface{})
	value, ok = clean[key]
	return value, ok
}

func (m *DeepCopyMap) Store(key, value interface{}) {
	m.mu.Lock()
	dirty := m.dirty()
	dirty[key] = value
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *DeepCopyMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	clean, _ := m.clean.Load().(map[interface{}]interface{})
	actual, loaded = clean[key]
	if loaded {
		return actual, loaded
	}

	m.mu.Lock()
	// Reload clean in case it changed while we were waiting on m.mu.
	clean, _ = m.clean.Load().(map[interface{}]interface{})
	actual, loaded = clean[key]
	if !loaded {
		dirty := m.dirty()
		dirty[key] = value
		actual = value
		m.clean.Store(dirty)
	}
	m.mu.Unlock()
	return actual, loaded
}

func (m *DeepCopyMap) Delete(key interface{}) {
	m.mu.Lock()
	dirty := m.dirty()
	delete(dirty, key)
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *DeepCopyMap) Range(f func(key, value interface{}) (shouldContinue bool)) {
	clean, _ := m.clean.Load().(map[interface{}]interface{})
	for k, v := range clean {
		if !f(k, v) {
			break
		}
	}
}

func (m *DeepCopyMap) dirty() map[interface{}]interface{} {
	clean, _ := m.clean.Load().(map[interface{}]interface{})
	dirty := make(map[interface{}]interface{}, len(clean)+1)
	for k, v := range clean {
		dirty[k] = v
	}
	return dirty
}

