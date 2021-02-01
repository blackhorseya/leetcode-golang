package problem_0007

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123 then 321",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "-123 then -321",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "120 then 21",
			args: args{120},
			want: 21,
		},
		{
			name: "0 then 0",
			args: args{x: 0},
			want: 0,
		},
		{
			name: "max int64 then 0",
			args: args{x: math.MaxInt64},
			want: 0,
		},
		{
			name: "min int64 then 0",
			args: args{x: math.MinInt64},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
