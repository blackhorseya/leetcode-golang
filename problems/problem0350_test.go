package problems

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2, 2},
		},
		{
			name: "example 2",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{4, 9},
		},
		{
			name: "no intersection",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6},
			},
			want: []int{},
		},
		{
			name: "empty nums1",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "empty nums2",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{},
			},
			want: []int{},
		},
		{
			name: "both empty",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: []int{},
		},
		{
			name: "one element intersection",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{3, 4, 5},
			},
			want: []int{3},
		},
		{
			name: "multiple identical elements",
			args: args{
				nums1: []int{1, 1, 1, 2, 2, 3},
				nums2: []int{1, 1, 2, 3, 3},
			},
			want: []int{1, 1, 2, 3},
		},
		{
			name: "nums1 larger than nums2",
			args: args{
				nums1: []int{1, 1, 1, 1, 1, 2, 2, 3},
				nums2: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect2(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
