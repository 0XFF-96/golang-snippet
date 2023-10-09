package _map

// This file contains reference map implementations for unit-tests.

// mapInterface is the interface Map implements.
type mapInterface interface {
	Load(interface{}) (interface{}, bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Delete(interface{})
	Range(func(key, value interface{}) (shouldContinue bool))
}
