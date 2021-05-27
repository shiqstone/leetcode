package code059

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				n: 1,
			},
			want: [][]int{
				[]int{1},
			},
		},
		{
			name: "test2",
			args: args{
				n: 2,
			},
			want: [][]int{
				[]int{1, 2},
				[]int{4, 3},
			},
		},
		{
			name: "test3",
			args: args{
				n: 3,
			},
			want: [][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
