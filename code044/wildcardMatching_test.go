package code044

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		input  string
		parten string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test0",
			args: args{
				input:  "",
				parten: "",
			},
			want: true,
		},
		{
			name: "test1",
			args: args{
				input:  "",
				parten: "?",
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				input:  "a",
				parten: "",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				input:  "ab",
				parten: "*",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				input:  "aa",
				parten: "a",
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				input:  "aa",
				parten: "a?",
			},
			want: true,
		},
		{
			name: "test6",
			args: args{
				input:  "cb",
				parten: "?a",
			},
			want: false,
		},
		{
			name: "test7",
			args: args{
				input:  "adceb",
				parten: "a*b",
			},
			want: true,
		},
		{
			name: "test8",
			args: args{
				input:  "acdcb",
				parten: "a*c?b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.input, tt.args.parten); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
