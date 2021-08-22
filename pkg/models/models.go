package models

import (
	"strconv"
	"strings"
)

// ListNode declare ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var ret []string
	for l != nil {
		ret = append(ret, strconv.Itoa(l.Val))
		l = l.Next
	}

	return strings.Join(ret, ",")
}

// NewListNodeFromInts create a list node from int slice
func NewListNodeFromInts(numbers []int) *ListNode {
	ret := &ListNode{}
	next := ret
	for _, number := range numbers {
		node := &ListNode{Val: number}
		next.Next = node
		next = node
	}

	return ret.Next
}
