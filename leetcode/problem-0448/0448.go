package problem_0448

func findDisappearedNumbers(nums []int) []int {
	l := make([]int, len(nums) + 1)
	for _, num := range nums {
		l[num] += 1
	}

	var ret []int
	for i, val := range l[1:] {
		if val == 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}
