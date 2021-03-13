package _map

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// go map 的性能问题
// https://github.com/golang/go/issues/20135

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func TestMap(t *testing.T) {
	var previous uint64 = 0

	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc at startup:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc

	elements := make([]string, 0)
	for i := 0; i < 200000; i++ {
		elements = append(elements, RandStringBytesMaskImprSrc(250))
	}

	stats = new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc Elements M1:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc

	m1 := make(map[string]string, 0)
	for _, v := range elements {
		m1[RandStringBytesMaskImprSrc(20)] = v
	}

	stats = new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc After M1:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc

	m2 := make(map[string]string, 0)
	for k, v := range m1 {
		m2[RandStringBytesMaskImprSrc(20)] = v
		delete(m1, k)
	}

	stats = new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc After M2:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc
}

