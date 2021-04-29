package main

import (
	"reflect"
	"testing"
)

func Test_searchPos(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{
				3, 4,
			},
		},
		{
			name: "test2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{
				-1, -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchPos(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
