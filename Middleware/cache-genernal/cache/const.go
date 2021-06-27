package cache

import (
	"time"
)

const (
	statInterval     = time.Minute


	defaultCacheName = "proc"

	// make the expiry unstable to avoid lots of cached items expire at the same time
	// make the unstable expiry to be [0.95, 1.05] * seconds
	expiryDeviation = 0.05


	slots            = 300
)
