package advance_concurrency

import (
	"time"
)

// Fetch fetches Items for uri and returns the time when the next
// fetch should be attempted.  On failure, Fetch returns an error.
func Fetch(domain string) Fetcher { return nil } // fetches Items from domain

type Item struct{
	Title, Channel, GUID string // a subset of RSS fields
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

type Subscription interface {
	Updates() <-chan Item // stream of Items
	Close() error         // shuts down the stream
}

// merges several streams
func Merge(subs ...Subscription) Subscription {return nil}

