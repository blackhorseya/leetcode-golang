package listnode

import (
	"reflect"
	"testing"
)

func TestNewFromSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1]",
			args: args{slice: []int{1}},
			want: &ListNode{
				Value: 1,
				Next:  nil,
			},
		},
		{
			name: "[1,2]",
			args: args{slice: []int{1, 2}},
			want: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 2,
					Next:  nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
