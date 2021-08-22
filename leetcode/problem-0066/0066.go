package problem

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i]/10 == 0 {
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
