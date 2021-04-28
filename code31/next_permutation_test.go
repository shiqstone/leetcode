package main

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "test2",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: []int{1, 5, 1},
		},
		{
			name: "test4",
			args: args{
				nums: []int{4, 1, 2, 3},
			},
			want: []int{4, 1, 3, 2},
		},
		{
			name: "test5",
			args: args{
				nums: []int{1, 5, 8, 4, 7, 6, 5, 3, 1},
			},
			want: []int{1, 5, 8, 5, 1, 3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
