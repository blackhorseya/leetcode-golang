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

func TestListNode_String(t *testing.T) {
	type fields struct {
		Value int
		Next  *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "[1,2,3]",
			fields: fields{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: nil}}},
			want:   "1->2->3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Value: tt.fields.Value,
				Next:  tt.fields.Next,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
