package code049

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				[]string{"eat", "tea", "ate"},
				[]string{"tan", "nat"},
				[]string{"bat"},
			},
		},
		{
			name: "test1",
			args: args{
				strs: []string{"ac", "c"},
			},
			want: [][]string{
				[]string{"ac"},
				[]string{"c"},
			},
		},
		{
			name: "test1",
			args: args{
				strs: []string{"ac", "d"},
			},
			want: [][]string{
				[]string{"ac"},
				[]string{"d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
