package problem_0053

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1] then 1",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "[5,4,-1,7,8] then 23",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
		{
			name: "[-2,1,-3,4,-1,2,1,-5,4] then 6",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
