package problems

func moveZeroes(nums []int) {
	current := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[current] = nums[i]
			current++
		}
	}

	for i := current; i < len(nums); i++ {
		nums[i] = 0
	}
}
