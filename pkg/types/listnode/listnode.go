package listnode

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
