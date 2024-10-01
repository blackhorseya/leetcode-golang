package problem0102

import (
	"container/list"

	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func levelOrder(root *TreeNode) [][]int {
	queue := list.New()
	queue.PushBack(root)
	res := make([][]int, 0)
	for queue.Len() > 0 {
		size := queue.Len()
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node == nil {
				continue
			}
			level = append(level, node.Val.(int))
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		}
		if len(level) > 0 {
			res = append(res, level)
		}
	}

	return res
}
