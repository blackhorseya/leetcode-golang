package problems

import (
	"testing"

	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "example 1",
		// 	args: args{head: ArrayToLinkedList([]int{1, 2, 2, 1})},
		// 	want: true,
		// },
		// {
		// 	name: "example 2",
		// 	args: args{head: ArrayToLinkedList([]int{1, 2})},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
