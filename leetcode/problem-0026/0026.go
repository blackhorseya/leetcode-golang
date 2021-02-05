package problem_0026

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	ret := 0
	for i := 0; i< len(nums) - 1; i++ {
		if nums[i] != nums[i+1] {
			ret++
			nums[ret] = nums[i+1]
		}
	}

	return ret+1
}
