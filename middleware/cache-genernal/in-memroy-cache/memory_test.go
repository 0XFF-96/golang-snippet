package in_memroy_cache

import (
	"testing"
	"sync"
)

type ringStripe struct {
	store    []uint64
	capacity uint64
}

func newRingStripe(capacity uint64) *ringStripe {
	return &ringStripe{
		store:    make([]uint64, 0, capacity),
		capacity: capacity,
	}
}

func (s *ringStripe) PushOrReturn(item uint64) []uint64 {
	s.store = append(s.store, item)
	if uint64(len(s.store)) >= s.capacity {
		ret := s.store[:]
		s.store = make([]uint64, 0, s.capacity)
		return ret
	}
	return nil
}

type ringBuffer struct {
	stripes []*ringStripe
	pool    *sync.Pool
}

func newRingBuffer(capacity uint64) *ringBuffer {
	return &ringBuffer{
		pool: &sync.Pool{
			New: func() interface{} {
				return newRingStripe(capacity)
			},
		},
	}
}

func (b *ringBuffer) PushOrReturn(item uint64) []uint64 {
	stripe := b.pool.Get().(*ringStripe)
	defer b.pool.Put(stripe)

	got := stripe.PushOrReturn(item)
	return got
}

func TestMemory(t *testing.T) {
	t.Log()
}

