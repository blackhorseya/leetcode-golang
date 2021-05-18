package problem_0070

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 then 2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "3 then 3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "5 then 8",
			args: args{n: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
