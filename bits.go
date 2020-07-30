package bits

// ref: https://github.com/morganxf/bit/blob/master/funcs.go#L19
// 二分法思想
// SetAllToOnesFromFirstOne64 最左非0bit后的bit全部置1
func SetAllToOnesFromFirstOne64(x uint64) uint64 {
	// x = 0x80000000
	// x = 1000 0000 | 0100 0000 = 1100 0000
	x |= x >> (1 >> 0) // 1
	// x = 1100 0000 | 0011 0000 = 1111 0000
	x |= x >> (1 << 1) // 2
	// x = 1111 0000 | 0000 1111 = 1111 1111
	x |= x >> (1 << 2) // 4
	// x = 0xf0 | 0x0f = 0xff
	x |= x >> (1 << 3) // 8
	// x = 0xff00 | 0x00ff = 0xffff
	x |= x >> (1 << 4) // 16
	// x = 0xffff0000 | 0x0000ffff = 0xffffffff
	x |= x >> (1 << 5) // 32
	return x
}

func SetAllToOnesFromFirstOne64_1(x uint64) uint64 {
	n := GetMinNumBits64(x)
	return x | ((1<<64 - 1) >> (64 - n))
}

// ref: math/bits.Len64
// 二分法思想 + 查找表思想
// 通过二分法思想及右移快速计算前56bit中1的位置。后续可以继续使用二分法来获取1的位置，但是通过查找表更快
// 不能直接使用查找表，因为2^64数据量太大
func GetMinNumBits64(x uint64) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	// 此时如果仍然存在1bit，那么一定存在于最右8bit中。
	// 此时可以通过查找表来计算剩余所需bit数，8bit需要空间256
	return n + int(len8tab[x])
}

// 二分法思想：完全二分法，不使用查找表
func getMinNumBits64_1(x uint64) (n int) {
	// [32, 16, 8, 4, 2, 1]
	for i := 5; i >= 0; i-- {
		shift := 1 << i
		if x >= 1<<shift {
			x >>= shift
			n += shift
		}
	}
	if x >= 1 {
		n += 1
	}
	return n
}
