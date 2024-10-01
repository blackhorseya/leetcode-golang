package problem0019

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// removeNthFromEnd is to remove the n-th node from the end of the list and return its head
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// create a dummy node
	dummy := &ListNode{Next: head}

	// create two pointers
	first, second := dummy, dummy

	// move the first pointer to the n-th node from the beginning
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// move the first pointer to the end, maintaining the gap
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// remove the n-th node
	second.Next = second.Next.Next

	return dummy.Next
}
