package problem0125

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}

		if !isAlphanumeric(s[j]) {
			j--
			continue
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(u uint8) bool {
	return (u >= 'a' && u <= 'z') || (u >= 'A' && u <= 'Z') || (u >= '0' && u <= '9')
}

func toLower(u uint8) uint8 {
	if u >= 'A' && u <= 'Z' {
		return u + 32
	}

	return u
}
