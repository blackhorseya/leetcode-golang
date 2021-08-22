package problem

import "math"

// Node ...
type Node struct {
	Value int
	Min   int
}

// MinStack ...
type MinStack struct {
	Nodes []*Node
}

// Constructor ...
func Constructor() MinStack {
	return MinStack{}
}

// Push ...
func (s *MinStack) Push(val int) {
	min := s.GetMin()
	if val < min {
		min = val
	}

	s.Nodes = append(s.Nodes, &Node{
		Value: val,
		Min:   min,
	})
}

// Pop ...
func (s *MinStack) Pop() {
	if len(s.Nodes) == 0 {
		return
	}

	s.Nodes = s.Nodes[:len(s.Nodes)-1]
}

// Top ...
func (s *MinStack) Top() int {
	if len(s.Nodes) == 0 {
		return 0
	}

	return s.Nodes[len(s.Nodes)-1].Value
}

// GetMin ...
func (s *MinStack) GetMin() int {
	if len(s.Nodes) == 0 {
		return math.MaxInt32
	}

	return s.Nodes[len(s.Nodes)-1].Min
}
