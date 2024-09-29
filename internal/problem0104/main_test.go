package problem0104

import (
	"testing"

	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{root: SliceToTree([]any{3, 9, 20, nil, nil, 15, 7})},
			want: 3,
		},
		{
			name: "example 2",
			args: args{root: SliceToTree([]any{1, nil, 2})},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
				PrintTree(tt.args.root)
			}
		})
	}
}
