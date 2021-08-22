package problem_0268

func missingNumber(nums []int) int {
	top := 1
	high := len(nums)
	bottom := len(nums)
	expected := (top + bottom) * high / 2

	actual := 0
	for _, num := range nums {
		actual += num
	}

	return expected - actual
}
