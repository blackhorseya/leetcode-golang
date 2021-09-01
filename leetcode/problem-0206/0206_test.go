package problem

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/leetcode-golang/pkg/models"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *models.ListNode
	}
	tests := []struct {
		name string
		args args
		want *models.ListNode
	}{
		{
			name: "[1,2,3,4,5] then [5,4,3,2,1]",
			args: args{head: models.NewListNodeFromInts([]int{1, 2, 3, 4, 5})},
			want: models.NewListNodeFromInts([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
