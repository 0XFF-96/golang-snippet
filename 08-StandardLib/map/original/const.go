package original

import (
	"unsafe"
)

const (
	// 一个桶里面最多可以装的键值对的个数，8对。
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits

	// 触发扩容操作的最大装载因子的临界值
	loadFactor = 6.5

	// 为了保持内联，键 和 值 的最大长度都是128字节，如果超过了128个字节，就存储它的指针
	maxKeySize   = 128
	maxValueSize = 128

	// 数据偏移应该是 bmap 的整数倍，但是需要正确的对齐。
	dataOffset = unsafe.Offsetof(struct {
		b bmap
		v int64
	}{}.v)

	// tophash 的一些值
	empty          = 0 // 没有键值对
	evacuatedEmpty = 1 // 没有键值对，并且桶内的键值被迁移走了。
	evacuatedX     = 2 // 键值对有效，并且已经迁移了一个表的前半段
	evacuatedY     = 3 // 键值对有效，并且已经迁移了一个表的后半段
	minTopHash     = 4 // 最小的 tophash

	// 标记
	iterator     = 1 // 当前桶的迭代子
	oldIterator  = 2 // 旧桶的迭代子
	hashWriting  = 4 // 一个goroutine正在写入map
	sameSizeGrow = 8 // 当前字典增长到新字典并且保持相同的大小

	// 迭代子检查桶ID的哨兵
	// noCheck = 1<<(8*sys.PtrSize) - 1
)

