package problem_0003

func lengthOfLongestSubstring(s string) int {
	asc := make([]int, 128)
	left := 0
	ret := 0

	for i, ch := range s {
		if asc[ch] > left {
			left = asc[ch]
		}

		if n := i - left + 1; n > ret {
			ret = n
		}

		asc[ch] = i + 1
	}

	return ret
}
