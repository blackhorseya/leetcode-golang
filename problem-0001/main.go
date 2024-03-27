package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		need := target - num
		ami, ok := m[need]
		if ok {
			return []int{ami, i}
		}

		m[num] = i
	}

	return []int{}
}
