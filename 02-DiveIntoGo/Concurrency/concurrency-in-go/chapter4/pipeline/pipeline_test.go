package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}

	for _, v := range multiply(add(multiply(ints, 2), 1), 3) {
		fmt.Println(v)
	}

	// 批处理
	// 流处理
	for _, v := range ints {
		fmt.Println((2 * (v*2 + 1)))
	}
}

func TestPipelineStreamOriented(t *testing.T) {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}

	// limits our ability to scale
	// limits the reuse of how we feed the pipeline
	// 有三个点需要注意⚠️
	// 1、
	// Each stage is receiving and emitting a discrete value,
	// and the memory footprint of our program is back down to only the size of
	// the pipeline’s input.
	// But we had to pull the pipeline down into the body of the for loop
	// and let the range do the heavy lifting of feeding our pipeline.
}