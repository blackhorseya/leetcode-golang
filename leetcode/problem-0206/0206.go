package problem

import "github.com/blackhorseya/leetcode-golang/pkg/models"

func reverseList(head *models.ListNode) *models.ListNode {
	var prev *models.ListNode

	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}

	return prev
}
