package client

import (
	"errors"
	"time"
)

const (
	_registerURL = "http://%s/discovery/register"
	_setURL      = "http://%s/discovery/set"
	_cancelURL   = "http://%s/discovery/cancel"
	_renewURL    = "http://%s/discovery/renew"
	_pollURL     = "http://%s/discovery/polls"
	_registerGap = 30 * time.Second

	_statusUP = "1"

	_appid = "infra.discovery"
)

var (
	_ Builder  = &Discovery{}
	_ Registry = &Discovery{}

	// ErrDuplication duplication treeid.
	ErrDuplication = errors.New("discovery: instance duplicate registration")
)
