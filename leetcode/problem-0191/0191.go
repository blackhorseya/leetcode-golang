package problem

import "fmt"

func hammingWeight(num uint32) int {
	str := fmt.Sprintf("%b", num)
	count := 0
	for _, ch := range str {
		if ch == '1' {
			count++
		}
	}

	return count
}
