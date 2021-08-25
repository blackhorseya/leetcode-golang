package problem

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2 10 then 1024",
			args: args{x: 2, n: 10},
			want: float64(1024),
		},
		{
			name: "2.1 3 then 9.261",
			args: args{x: 2.1, n: 3},
			want: 9.261,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
