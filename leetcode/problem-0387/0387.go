package problem

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, ch := range s {
		m[int(ch-'a')]++
	}

	for i, ch := range s {
		if m[int(ch-'a')] == 1 {
			return i
		}
	}

	return -1
}
