package problem

func search(nums []int, target int) int {
	ret := -1
	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return ret
}
