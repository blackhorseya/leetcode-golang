package problem0102

import (
	"reflect"
	"testing"

	. "github.com/blackhorseya/leetcode-golang/pkg"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{root: SliceToTree([]any{3, 9, 20, nil, nil, 15, 7})},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "test case 2",
			args: args{root: SliceToTree([]any{1})},
			want: [][]int{{1}},
		},
		{
			name: "test case 3",
			args: args{root: SliceToTree([]any{})},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
