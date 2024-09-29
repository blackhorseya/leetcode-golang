package problem0101

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// isSymmetric represent for is symmetric tree
// Time complexity: O(n)
// Space complexity: O(n)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
