package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1,2,3,4,5] then [5,4,3,2,1]",
			args: args{head: NewListNode([]int{1, 2, 3, 4, 5})},
			want: NewListNode([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %s, want %s", got, tt.want)
			}
		})
	}
}
