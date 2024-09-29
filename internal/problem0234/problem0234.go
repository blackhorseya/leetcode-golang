package problem0234

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// isPalindrome represent for LeetCode 234
// Time complexity: O(n)
// Space complexity: O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// find the middle of the linked list by using the slow and fast pointer
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half of the linked list
	secondHalf := reverse(slow)

	// compare the first half and the second half of the linked list
	for secondHalf != nil {
		if head.Val != secondHalf.Val {
			return false
		}
		head = head.Next
		secondHalf = secondHalf.Next
	}

	return true
}

// reverse the linked list
// Time complexity: O(n)
// Space complexity: O(1)
func reverse(next *ListNode) *ListNode {
	var prev *ListNode

	for next != nil {
		temp := next.Next
		next.Next = prev
		prev = next
		next = temp
	}

	return prev
}
