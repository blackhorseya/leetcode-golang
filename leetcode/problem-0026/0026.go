package problem

func removeDuplicates(nums []int) int {
	ret := 1
	for i := 1; i < len(nums); i++ {
		if nums[ret-1] != nums[i] {
			nums[ret] = nums[i]
			ret++
		}
	}

	return ret
}
