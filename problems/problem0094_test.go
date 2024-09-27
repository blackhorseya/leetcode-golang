package problems

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{root: SliceToTree([]any{1, nil, 2, nil, nil, 3, nil})},
			want: []int{1, 3, 2},
		},
		{
			name: "example 2",
			args: args{root: SliceToTree([]any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, nil, nil, 9, nil})},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name: "example 3",
			args: args{root: SliceToTree([]any{})},
			want: []int{},
		},
		{
			name: "example 4",
			args: args{root: SliceToTree([]any{1})},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintTree(tt.args.root)
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
