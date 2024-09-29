package problem0094

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	inOrderHelper(root, &res)
	return res
}

func inOrderHelper(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	inOrderHelper(node.Left, res)
	*res = append(*res, node.Val.(int))
	inOrderHelper(node.Right, res)
}
