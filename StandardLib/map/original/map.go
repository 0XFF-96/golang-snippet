package original

import (
	"unsafe"
)

type hmap struct {
	count     int // map 的长度
	flags     uint8
	B         uint8  // log以2为底，桶个数的对数 (总共能存 6.5 * 2^B 个元素)
	noverflow uint16 // 近似溢出桶的个数
	hash0     uint32 // 哈希种子

	buckets    unsafe.Pointer // 有 2^B 个桶的数组. count==0 的时候，这个数组为 nil.
	oldbuckets unsafe.Pointer // 旧的桶数组一半的元素
	nevacuate  uintptr        // 扩容增长过程中的计数器

	extra *mapextra // 可选字段
}

type mapextra struct {
	overflow [2]*[]*bmap
	nextOverflow *bmap
}

type bmap struct {
	tophash [bucketCnt]uint8
}
