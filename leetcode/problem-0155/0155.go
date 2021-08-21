package problem_0155

import "math"

type Node struct {
	Value int
	Min   int
}

type MinStack struct {
	Nodes []*Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if val < min {
		min = val
	}

	this.Nodes = append(this.Nodes, &Node{
		Value: val,
		Min:   min,
	})
}

func (this *MinStack) Pop() {
	if len(this.Nodes) == 0 {
		return
	}

	this.Nodes = this.Nodes[:len(this.Nodes)-1]
}

func (this *MinStack) Top() int {
	if len(this.Nodes) == 0 {
		return 0
	}

	return this.Nodes[len(this.Nodes)-1].Value
}

func (this *MinStack) GetMin() int {
	if len(this.Nodes) == 0 {
		return math.MaxInt32
	}

	return this.Nodes[len(this.Nodes)-1].Min
}
