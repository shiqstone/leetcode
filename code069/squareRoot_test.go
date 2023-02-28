package code069

import "testing"

func Test_sqrtRoot(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "sqrt 4",
			args: args{
				x: 4,
			},
			wantRes: 2,
		},
		{
			name: "sqrt 8",
			args: args{
				x: 8,
			},
			wantRes: 2,
		},
		{
			name: "sqrt 16",
			args: args{
				x: 16,
			},
			wantRes: 4,
		},
		{
			name: "sqrt 25",
			args: args{
				x: 25,
			},
			wantRes: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sqrtRoot(tt.args.x); gotRes != tt.wantRes {
				t.Errorf("sqrtRoot() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
