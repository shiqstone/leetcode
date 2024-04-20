package code085

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				matrix: [][]byte{{'0'}},
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				matrix: [][]byte{{'1'}},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				matrix: [][]byte{{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'}},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
