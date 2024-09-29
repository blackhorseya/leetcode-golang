package problem0108

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "example 1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			want: SliceToTree([]any{0, -10, 5, nil, -3, nil, 9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
				fmt.Println("got:")
				PrintTree(got)

				fmt.Println("want:")
				PrintTree(tt.want)
			}
		})
	}
}
