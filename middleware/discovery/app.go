package discovery

import (
	"sync"
)

// App Instances distinguished by hostname
type App struct {
	AppID           string
	Zone            string
	instances       map[string]*Instance
	latestTimestamp int64

	lock sync.RWMutex
}

