package problem

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	ret, max := 0, 0
	for k, v := range m {
		if v > max {
			max = v
			ret = k
		}
	}

	return ret
}
