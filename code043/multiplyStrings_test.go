package code043

import (
	"testing"
)

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 x 2",
			args: args{
				num1: "1",
				num2: "2",
			},
			want: "2",
		},
		{
			name: "12 x 12",
			args: args{
				num1: "12",
				num2: "12",
			},
			want: "144",
		},
		{
			name: "123 x 456",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "999 x 999",
			args: args{
				num1: "999",
				num2: "999",
			},
			want: "998001",
		},
		{
			name: "123456789 x 987654321",
			args: args{
				num1: "123456789",
				num2: "987654321",
			},
			want: "121932631112635269",
		},
		{
			name: "584 x 18",
			args: args{
				num1: "584",
				num2: "18",
			},
			want: "10512",
		},
		{
			name: "79362 x 217",
			args: args{
				num1: "79362",
				num2: "217",
			},
			want: "17221554",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 + 9",
			args: args{
				a: "1",
				b: "9",
			},
			want: "10",
		},
		{
			name: "72 + 1440",
			args: args{
				a: "72",
				b: "1440",
			},
			want: "1512",
		},
		{
			name: "9 + 10001",
			args: args{
				a: "9",
				b: "10001",
			},
			want: "10010",
		},
		{
			name: "78554 + 1953000",
			args: args{
				a: "78554",
				b: "1953000",
			},
			want: "2031554",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumStr(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
