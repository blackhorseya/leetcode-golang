package problem_0171

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A then 1",
			args: args{s: "A"},
			want: 1,
		},
		{
			name: "AB then 28",
			args: args{s: "AB"},
			want: 28,
		},
		{
			name: "ZY then 701",
			args: args{s: "ZY"},
			want: 701,
		},
		{
			name: "FXSHRXW then 2147483647",
			args: args{s: "FXSHRXW"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
