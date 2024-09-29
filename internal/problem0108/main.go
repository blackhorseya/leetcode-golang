package problem0108

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// sortedArrayToBST represents the method to convert sorted array to binary search tree
// Time complexity: O(n)
// Space complexity: O(log n)
func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := left + (right-left)/2
		root := &TreeNode{Val: nums[mid]}
		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)

		return root
	}

	return dfs(0, len(nums)-1)
}
