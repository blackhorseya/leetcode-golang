package leetcode

import "math"

func reverse(x int) int {
	ret := 0

	for value := x; value != 0; value /= 10 {
		ret = ret*10 + value%10
	}

	if ret < math.MinInt32 || ret > math.MaxInt32 {
		return 0
	}

	return ret
}
