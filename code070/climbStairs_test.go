package code070

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
			name: "test2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "test4",
			args: args{
				n: 4,
			},
			want: 5,
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
