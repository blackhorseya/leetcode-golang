package problem

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "\"\" \"\" then 0",
			args: args{haystack: "", needle: ""},
			want: 0,
		},
		{
			name: "aaaaa bba then -1",
			args: args{haystack: "aaaaa", needle: "bba"},
			want: -1,
		},
		{
			name: "hello ll then 2",
			args: args{haystack: "hello", needle: "ll"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
