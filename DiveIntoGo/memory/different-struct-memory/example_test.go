package different_struct_memory

import (
	"fmt"
	"testing"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	// 8
	a bool
	e byte
	c int8
	b int32

	// 8
	d int64
}

// 参考文档
func TestM(t *testing.T) {
	part1 := Part1{}

	part2 := Part2{}
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))

	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
}