package problem_0026

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1] then 1 [1]",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "[1,1,2] then 2 [1,2]",
			args: args{[]int{1, 1, 2}},
			want: 2,
		},
		{
			name: "[0,0,1,1,1,2,2,3,3,4] then 5 [0,1,2,3,4]",
			args: args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
