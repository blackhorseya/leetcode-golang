package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for index, num := range nums {
		if pos, ok := m[target-num]; ok {
			return []int{pos, index}
		}

		m[num] = index
	}

	return []int{}
}
