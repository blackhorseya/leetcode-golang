package problem0160

import (
	. "github.com/blackhorseya/leetcode-golang/pkg"
)

// getIntersectionNode 暴力解法
// Time complexity: O(m*n)
// Space complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for pa := headA; pa != nil; pa = pa.Next {
		for pb := headB; pb != nil; pb = pb.Next {
			if pa == pb {
				return pa
			}
		}
	}

	return nil
}

// getIntersectionNode2 hash set
// Time complexity: O(m+n)
// Space complexity: O(m)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	set := make(map[*ListNode]struct{})

	for pa := headA; pa != nil; pa = pa.Next {
		set[pa] = struct{}{}
	}

	for pb := headB; pb != nil; pb = pb.Next {
		if _, ok := set[pb]; ok {
			return pb
		}
	}

	return nil
}

// getIntersectionNode3 two pointer
// Time complexity: O(m+n)
// Space complexity: O(1)
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB

	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}

	return pa
}
