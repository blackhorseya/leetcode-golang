package problem0104

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// maxDepth is to get the maximum depth of a binary tree
// Time complexity: O(n)
// Space complexity: O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
