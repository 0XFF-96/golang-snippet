package discovery

import (
	"sync"
)

// Apps app distinguished by zone
type Apps struct {
	apps            map[string]*App
	lock            sync.RWMutex
	latestTimestamp int64
}


// NewApp new App.
func NewApp(zone, appid string) (a *App) {
	return
}

// Instances return slice of instances.
func (a *App) Instances() (is []*Instance) {
	return
}

// NewInstance new a instance.
func (a *App) NewInstance(ni *Instance, latestTime int64) (i *Instance, ok bool) {
	return
}

// Renew new a instance.
func (a *App) Renew(hostname string) (i *Instance, ok bool) {
	return
}

func (a *App) updateLatest(latestTime int64) {
}

// Cancel cancel a instance.
func (a *App) Cancel(hostname string, latestTime int64) (i *Instance, l int, ok bool) {
	return
}

// Len returns the length of instances.
func (a *App) Len() (l int) {
	return 0
}

// Set set new status,metadata,color  of instance .
func (a *App) Set(changes *ArgSet) (ok bool) {
	return false
}


