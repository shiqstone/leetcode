package code38

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "test2",
			args: args{
				n: 3,
			},
			want: "21",
		},
		{
			name: "test3",
			args: args{
				n: 4,
			},
			want: "1211",
		},
		{
			name: "test4",
			args: args{
				n: 5,
			},
			want: "111221",
		},
		{
			name: "test5",
			args: args{
				n: 6,
			},
			want: "312211",
		},
		{
			name: "test6",
			args: args{
				n: 7,
			},
			want: "13112221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay2(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
