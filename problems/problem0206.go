package problems

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
