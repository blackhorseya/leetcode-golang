package problem

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 0 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		k, j := i, i
		for k < len(s)-1 && s[k] == s[k+1] {
			k++
		}

		i = k + 1
		for k < len(s)-1 && j > 0 && s[k+1] == s[j-1] {
			k++
			j--
		}
		newLen := k - j + 1
		if newLen > maxLen {
			maxLen = newLen
			start = j
		}
	}

	return s[start : start+maxLen]
}
