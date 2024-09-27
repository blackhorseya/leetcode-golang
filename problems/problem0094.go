package problems

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	inOrder(root, &res)
	return res
}

func inOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, res)
	*res = append(*res, node.Val.(int))
	inOrder(node.Right, res)
}
