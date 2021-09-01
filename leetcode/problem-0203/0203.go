package problem

import "github.com/blackhorseya/leetcode-golang/pkg/models"

func removeElements(head *models.ListNode, val int) *models.ListNode {
	if head == nil {
		return head
	}

	cur := head
	for ;cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}