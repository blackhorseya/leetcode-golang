package problem

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] >= 2 {
			return true
		}
	}

	return false
}
