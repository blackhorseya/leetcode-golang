package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var ret *ListNode
	for _, num := range nums {
		ret = ret.Append(num)
	}

	return ret
}

func (root *ListNode) String() string {
	var vals []string

	temp := root
	for temp != nil {
		vals = append(vals, strconv.Itoa(temp.Val))
		temp = temp.Next
	}

	return strings.Join(vals, "->")
}

func (root *ListNode) Append(val int) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	if root == nil {
		return node
	}

	ptr := root
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = node

	return root
}

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = rev
		rev = head
		head = temp
	}

	return rev
}
