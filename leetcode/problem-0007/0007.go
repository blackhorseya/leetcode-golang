package problem_0007

func reverse(x int) int {
	sig, val, ret := 1, x, 0
	if val < 0 {
		sig = -1
		val = 0 - x
	}

	for val > 0 {
		ret = ret * 10 + val % 10
		val = val / 10
	}

	if ret >= 1 << 31 {
		return 0
	}

	return ret * sig
}
