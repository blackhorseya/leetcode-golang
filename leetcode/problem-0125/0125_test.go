package problem

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "\"race a car\" then false",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "\"A man, a plan, a canal: Panama\" then true",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
