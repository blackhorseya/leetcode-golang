package problem

func isPowerOfTwo(n int) bool {
	for n != 1 {
		if n%2 > 0 || n <= 0 {
			break
		}

		n /= 2
	}

	return n == 1
}
