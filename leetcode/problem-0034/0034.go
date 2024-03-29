package problem

func searchRange(nums []int, target int) []int {
	left, right := -1, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target && left == -1 {
			left = i
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target && right == -1 {
			right = i
		}
	}

	return []int{left, right}
}
