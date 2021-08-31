package problem

func isVowel(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}

	if ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}

	return false
}

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	ret := []byte(s)

	for left < right {
		if !isVowel(ret[left]) {
			left++
		}

		if !isVowel(ret[right]) {
			right--
		}

		if isVowel(s[left]) && isVowel(s[right]) {
			ret[left], ret[right] = ret[right], ret[left]
			left++
			right--
		}
	}

	return string(ret)
}
