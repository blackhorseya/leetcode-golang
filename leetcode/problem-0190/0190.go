package problem_0190

func reverseBits(num uint32) uint32 {
	var (
		ret    = uint32(0)
		mask   = uint32(1)
		length = 32
	)

	for i := 0; i < length; i++ {
		pop := (num >> i) & mask
		ret = ret | (pop << (length - i - 1))
	}

	return ret
}
