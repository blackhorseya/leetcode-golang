package problem_0172

func trailingZeroes(n int) int {
	ret :=0
	for n >= 5 {
		n /= 5
		ret += n
	}

	return ret
}
