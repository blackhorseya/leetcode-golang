package listnode

import (
	"fmt"
	"strings"
)

// ListNode is a node of linked list
type ListNode struct {
	Value int
	Next  *ListNode
}

// NewFromSlice creates a linked list from a slice
func NewFromSlice(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	for _, val := range slice {
		return &ListNode{
			Value: val,
			Next:  NewFromSlice(slice[1:]),
		}
	}

	return nil
}

func (l *ListNode) String() string {
	slice := make([]string, 0)
	current := l
	for current != nil {
		slice = append(slice, fmt.Sprintf("%d", current.Value))
		current = current.Next
	}

	return strings.Join(slice, "->")
}

func (l *ListNode) Append(val int) {
	current := l
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &ListNode{
		Value: val,
		Next:  nil,
	}
}
