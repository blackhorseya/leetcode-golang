package problem

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/leetcode-golang/pkg/models"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *models.ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *models.ListNode
	}{
		{
			name: "[1,2,6,3,4,5,6] 6 then [1,2,3,4,5]",
			args: args{head: models.NewListNodeFromInts([]int{1, 2, 6, 3, 4, 5, 6}), val: 6},
			want: models.NewListNodeFromInts([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "[7,7,7,7] 7 then []",
			args: args{head: models.NewListNodeFromInts([]int{7, 7, 7, 7}), val: 7},
			want: models.NewListNodeFromInts([]int{}),
		},
		{
			name: "[] 1 then []",
			args: args{head: models.NewListNodeFromInts([]int{}), val: 1},
			want: models.NewListNodeFromInts([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
