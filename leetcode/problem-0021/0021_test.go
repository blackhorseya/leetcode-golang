package problem

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/leetcode-golang/pkg/models"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *models.ListNode
		l2 *models.ListNode
	}
	tests := []struct {
		name string
		args args
		want *models.ListNode
	}{
		{
			name: "nil nil then nil",
			args: args{l1: nil, l2: nil},
			want: nil,
		},
		{
			name: "nil [0] then [0]",
			args: args{l1: nil, l2: models.NewListNodeFromInts([]int{0})},
			want: models.NewListNodeFromInts([]int{0}),
		},
		{
			name: "[1,2,4] [1,3,4] then [1,1,2,3,4,4]",
			args: args{l1: models.NewListNodeFromInts([]int{1, 2, 4}), l2: models.NewListNodeFromInts([]int{1, 3, 4})},
			want: models.NewListNodeFromInts([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
