package problem_0283

func moveZeroes(nums []int) {
	lastNonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}
