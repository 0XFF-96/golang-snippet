package go_zero_rollingwindow

import (
	"sync"
	"time"
)

type RollingWindow struct {
	lock          sync.RWMutex
	size          int
	win           *window
	interval      time.Duration
	offset        int
	ignoreCurrent bool
	lastTime      time.Duration
}


type window struct {
	buckets []*Bucket
	size    int
}
