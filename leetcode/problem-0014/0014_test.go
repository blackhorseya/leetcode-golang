package problem_0014

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[] then empty",
			args: args{[]string{}},
			want: "",
		},
		{
			name: "[flag, , flow] then empty",
			args: args{[]string{"flag", "", "flow"}},
			want: "",
		},
		{
			name: "[\"flower\",\"flow\",\"flight\"] then fl",
			args: args{[]string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "[\"dog\",\"racecar\",\"car\"] then empty",
			args: args{[]string{"dog", "racecar", "car"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
