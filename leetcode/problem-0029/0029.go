package problem

import "math"

func divide(dividend int, divisor int) int {
	ret := dividend / divisor

	if ret > math.MaxInt32 {
		return math.MaxInt32
	}

	if ret < math.MinInt32 {
		return math.MinInt32
	}

	return ret
}
