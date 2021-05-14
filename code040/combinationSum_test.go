package code040

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				candidates: []int{10, 1, 2, 6, 7, 1, 5},
				target:     8,
			},
			want: [][]int{
				[]int{1, 7},
				[]int{1, 2, 5},
				[]int{2, 6},
				[]int{1, 1, 6},
			},
		},
		{
			name: "test2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{
				[]int{1, 2, 2},
				[]int{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
