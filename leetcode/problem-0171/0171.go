package problem_0171

func titleToNumber(s string) int {
	col := 0
	for _, ch := range s {
		col *= 26
		col += int(ch - 'A') + 1
	}

	return col
}
