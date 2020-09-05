package mock_interface

import (
	"unsafe"
)

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter uintptr
	_type uintptr
	link uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

type eface struct {
	_type uintptr
	data unsafe.Pointer
}
