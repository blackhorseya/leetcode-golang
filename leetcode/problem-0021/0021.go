package problem_0021

import (
	"github.com/blackhorseya/leetcode-golang/pkg/models"
)

func mergeTwoLists(l1 *models.ListNode, l2 *models.ListNode) *models.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	ret := &models.ListNode{}
	next := ret
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			next.Next = l2
			l2 = l2.Next
		} else {
			next.Next = l1
			l1 = l1.Next
		}

		next = next.Next
	}

	if l1 == nil {
		next.Next = l2
	}

	if l2 == nil {
		next.Next = l1
	}

	return ret.Next
}
