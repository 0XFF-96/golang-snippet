package msyql

import (
	"sync"

	"source-code-reading/go/src/database/sql/driver"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]driver.Driver)
)


