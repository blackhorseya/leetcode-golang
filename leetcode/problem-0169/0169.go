package problem_0169

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
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
