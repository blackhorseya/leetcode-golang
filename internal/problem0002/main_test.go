package problem0002

import (
	"reflect"
	"testing"

	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				l1: ArrayToLinkedList([]int{2, 4, 3}),
				l2: ArrayToLinkedList([]int{5, 6, 4}),
			},
			want: ArrayToLinkedList([]int{7, 0, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
