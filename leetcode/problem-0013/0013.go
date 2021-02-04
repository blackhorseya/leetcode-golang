package problem_0013

var (
	lookup = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func romanToInt(s string) int {
	ret := 0
	for i, c := range s {
		val := lookup[string(c)]
		if i == 0 {
			ret += val
			continue
		}

		last := lookup[string(s[i-1])]
		if last < val {
			ret -= last * 2
		}

		ret += val
	}

	return ret
}
