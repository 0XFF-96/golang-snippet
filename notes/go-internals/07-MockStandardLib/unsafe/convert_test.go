package unsafe

import (
	"fmt"
	"testing"
)

// 类型转换
// It is important to remember that conversion keeps the integrity of the data converted.
// Nowadays, most of the languages use implicit / explicit conversion rather than simple cast.
func TestConvert(t *testing.T) {
// https://medium.com/a-journey-with-go/go-cast-vs-conversion-by-example-26e0ef3003f0
	var a int8 = -127

	// https://stackoverflow.com/questions/13870845/converting-from-an-integer-to-its-binary-representation
	fmt.Printf("cast uint8: %d\n", uint8(a))
	fmt.Printf("cast int16: %d\n", int16(a))

	// https://en.wikipedia.org/wiki/Signed_number_representations
	// https://medium.com/a-journey-with-go/go-cast-vs-conversion-by-example-26e0ef3003f0
	fmt.Printf("Bit pattern int8: %b\n", a)
	fmt.Printf("Bit pattern uint8: %b\n", uint8(a))
	fmt.Printf("Bit pattern int16: %b\n", int16(a))

	// --------
	// 内存布局
	a8 := fmt.Sprintf("%.32b", a)
	fmt.Println("a8:", a8)
	a8 = fmt.Sprintf("%.8b", a)
	fmt.Println("a8:", a8)
	a8 = fmt.Sprintf("%.8b", uint8(a))
	fmt.Println("a8:", a8)
	// -------
	var b int64 = 520
	fmt.Println(float64(b), float32(b), uint64(b), uint32(b), uint16(b))

	var c int64 = -520
	// 后面两个都 overflow 了
	fmt.Println(float64(c), float32(c), uint64(c), uint32(c), uint16(c))
}
