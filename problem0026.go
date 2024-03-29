package leetcode_golang

func removeDuplicates(nums []int) int {
	k := 1

	for i, num := range nums {
		if i == 0 {
			continue
		}

		if num == nums[k-1] {
			continue
		} else {
			nums[k] = num
			k++
		}
	}

	return k
}
