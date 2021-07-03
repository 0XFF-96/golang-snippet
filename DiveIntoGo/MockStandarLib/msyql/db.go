package msyql

import (
	"context"

	"source-code-reading/go/src/database/sql/driver"
)

func OpenDB(c driver.Connector) *DB {
	ctx, cancel := context.WithCancel(context.Background())
	db := &DB{
		connector:    c,
		openerCh:     make(chan struct{}, 10),
		resetterCh:   make(chan *driverConn, 50),
		lastPut:      make(map[*driverConn]string),
		connRequests: make(map[uint64]chan connRequest),
		stop:         cancel,
	}

	go db.connectionOpener(ctx)
	go db.connectionResetter(ctx)

	return db
}

// Runs in a separate goroutine, opens new connections when requested.
func (db *DB) connectionOpener(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-db.openerCh:
			// db.openNewConnection(ctx)
		}
	}
}

// connectionResetter runs in a separate goroutine to reset connections async
// to exported API.
func (db *DB) connectionResetter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(db.resetterCh)
			for dc := range db.resetterCh {
				dc.Unlock()
			}
			return
		case dc := <-db.resetterCh:
			dc.resetSession(ctx)
		}
	}
}
