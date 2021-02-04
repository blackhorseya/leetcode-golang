package models

import (
	"reflect"
	"testing"
)

func TestNewListNodeFromInts(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1] then [1]",
			args: args{numbers: []int{1}},
			want: &ListNode{Val: 1},
		},
		{
			name: "[1,2] then [1,2]",
			args: args{numbers: []int{1, 2}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "[2,3,5] then [2,3,5]",
			args: args{numbers: []int{2, 3, 5}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNodeFromInts(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNodeFromInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "[1,2,3]",
			fields: fields{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want:   "1,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
