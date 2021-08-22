package problem

import "fmt"

func countBits(n int) []int {
	ret := make([]int, n+1)
	for n >= 0 {
		count := 0
		for _, val := range fmt.Sprintf("%b", n) {
			if val == '1' {
				count++
			}
		}

		ret[n] = count

		n--
	}

	return ret
}
