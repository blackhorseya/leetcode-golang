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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ret *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			ret = ret.Append(list2.Val)
			list2 = list2.Next
		} else {
			ret = ret.Append(list1.Val)
			list1 = list1.Next
		}
	}

	for list1 != nil {
		ret = ret.Append(list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		ret = ret.Append(list2.Val)
		list2 = list2.Next
	}

	return ret
}
