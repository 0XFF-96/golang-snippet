package chapter2_float

import (
	"fmt"
	"math"
	"testing"
)

// GO 采用 IEEE754 浮点数标准

// QA1: 单精度和双精的区别？
//
func TestFloatAdd(t *testing.T) {
	var f1 float64 = 0.3
	var f2 float64 = 0.6
	fmt.Println(f1 + f2)  // the result is ? and why
}

func TestVisualizeFloat(t *testing.T) {
	var number float32 = 0.085
	fmt.Printf("Starting Number :%f \n\n", number)
	bits := math.Float32bits(number)
	binary := fmt.Sprintf("%.32b", bits)
	fmt.Printf("Bit pattern: %s \n", binary)

	// binary to represent value
	bias := 127
	sign := bits & ( 1 << 31)
	exponentRaw := int(bits >> 23)
	exponent := exponentRaw - bias

	var mantissa float64  // 尾数；假数；定点部分；小数部分
	// 遍历位数
	for index, bit := range binary[9:23]{

		// bit = 48 or bit = 49 ?
		// 0 或者是 1 的 ascii 码
		// bit == 49
		if string(bit) == "1" {
			position := index + 1
			bitValue := math.Pow(2, float64(position))
			fractional := 1 / bitValue

			// 这里会损失精度？
			mantissa = mantissa + fractional
		}
	}
	value := (1 + mantissa) * math.Pow(2, float64(exponent))
	fmt.Printf("Sign:%d Exponent: %d (%d) Mantissa: %f Value: %f \n",
		sign,
		exponentRaw,
		exponent,
		mantissa,
		value,
		)
}

func TestIsInt(t *testing.T) {
	// isInt()
}

// 浮点数格式，整数
// 1. 浮点数格式中指数的值大于 127 (0.1123)
// 2.

func isInt(bits uint32, bias int) bool {
	exponent := int(bits >> 23 ) - bias - 23

	// 通过位运算，计算小数部分数值
	coefficient := (bits & (( 1 << 23 ) - 1 )) | ( 1 << 23)
	intTest := coefficient & (1 << uint32(-exponent) - 1)
	fmt.Printf("\n Exponent:%d coefficient:%d intTest:%d \n",
		exponent,
		coefficient,
		intTest,
		)
	if exponent < -23 {
		return false
	}
	if exponent < 0 && intTest != 0 {
		return false
	}
	return true
}

// QA2: 什么是常规数和非常规数
// math.big 库


