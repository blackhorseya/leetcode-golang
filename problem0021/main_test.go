package main

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1,2,3]",
			args: args{nums: []int{1, 2, 3}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[1,2,3]",
			args: args{nums: []int{1, 2, 3}},
			want: "1->2->3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewListNode(tt.args.nums)

			if got := node.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1,2,4] [1,3,4] then [1,1,2,3,4,4]",
			args: args{list1: NewListNode([]int{1, 2, 4}), list2: NewListNode([]int{1, 3, 4})},
			want: NewListNode([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %s, want %s", got, tt.want)
			}
		})
	}
}
