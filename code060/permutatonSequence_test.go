package code060

import (
	"testing"
)

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				n: 3,
				k: 3,
			},
			want: "213",
		},
		{
			name: "test2",
			args: args{
				n: 4,
				k: 9,
			},
			want: "2314",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
