package problem

import "math"

var m = map[byte]int64{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func myAtoi(s string) int {
	ret, factor := int64(0), int64(1)
	i, j := 0, len(s)

	for i < j && s[i] == ' ' {
		i++
	}

	if i < j && s[i] == '-' {
		factor = -1
		i++
	} else if i < j && s[i] == '+' {
		i++
	}

	for i < j && ret < math.MaxInt32 {
		if val, ok := m[s[i]]; ok {
			ret = ret*10 + val
			i++
		} else {
			break
		}
	}

	ret = ret * factor

	if ret > math.MaxInt32 {
		return math.MaxInt32
	}

	if ret < math.MinInt32 {
		return math.MinInt32
	}

	return int(ret)
}
