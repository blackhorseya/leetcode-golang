package problem

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "\"42\" then 42",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "\"   -42\" then -42",
			args: args{s: "   -42"},
			want: -42,
		},
		{
			name: "\"4193 with words\" then 4193",
			args: args{s: "4193 with words"},
			want: 4193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
