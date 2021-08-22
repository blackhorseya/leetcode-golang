package problem

func singleNumber(nums []int) int {
	ret := nums[0]
	// all numbers execute xor operation, then same number is zero
	for _, num := range nums[1:] {
		ret = ret ^ num
	}

	return ret
}
