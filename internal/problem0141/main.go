package problem0141

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// hasCycle is to solve the problem
// Time complexity: O(n)
// Space complexity: O(n)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}

		m[head] = struct{}{}
		head = head.Next
	}

	return false
}

// hasCycle2 two pointers
// Time complexity: O(n)
// Space complexity: O(1)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
