package problem

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := max
	for _, val := range nums[1:] {
		if sum+val > val {
			sum = sum + val
		} else {
			sum = val
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
