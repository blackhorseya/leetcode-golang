package problem0002

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// addTwoNumbers represents the method to add two numbers
// Time complexity: O(max(m, n))
// Space complexity: O(max(m, n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry % 10}
		cur = cur.Next
		carry /= 10
	}

	return dummy.Next
}
