package problem

func climbStairs(n int) int {
	a, b := 1, 1

	for n > 0 {
		b += a
		a = b - a

		n--
	}

	return a
}
