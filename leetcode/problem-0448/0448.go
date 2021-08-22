package problem_0448

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

func findDisappearedNumbers(nums []int) []int {
	var ret []int
	want := len(nums)
	for i := 1; i <= want; i++ {
		if !contains(nums, i) {
			ret = append(ret, i)
		}
	}

	return ret
}
