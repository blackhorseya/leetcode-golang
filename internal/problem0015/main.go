package problem0015

import (
	"sort"
)

// threeSum 回溯演算法
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	var backtrack func(state int, choices []int, target int)
	backtrack = func(state int, choices []int, target int) {
		// 递归终止条件
		if len(choices) == 3 {
			if target == 0 {
				res = append(res, append([]int{}, choices...))
			}
			return
		}

		for i := state; i < len(nums); i++ {
			// 剪枝：跳過重複的數字
			if i > state && nums[i] == nums[i-1] {
				continue
			}

			// 剪枝：如果 nums[i] 大於 target，則後面的數字都會大於 target
			if len(choices) == 2 && nums[i] > target {
				break
			}

			// 嘗試選擇 nums[i]
			choices = append(choices, nums[i])

			// 遞歸
			backtrack(i+1, choices, target-nums[i])

			// 回溯
			choices = choices[:len(choices)-1]
		}
	}

	backtrack(0, []int{}, 0)

	return res
}
