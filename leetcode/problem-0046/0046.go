package problem

func permute(nums []int) [][]int {
	var ret [][]int

	dfs([]int{}, nums, &ret)

	return ret
}

func dfs(currComb, left []int, res *[][]int) {
	if 0 == len(left) {
		*res = append(*res, currComb)
		return
	}
	for idx, l := range left {
		dfs(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...),
			res,
		)
	}
}
