package code009

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para int
	ans  bool
}{
	{
		123,
		false,
	},
	{
		-1234,
		false,
	},
	{
		121,
		true,
	},
	{
		12321,
		true,
	},
	{
		12345654321,
		true,
	},
	{
		1234567899,
		false,
	},
}

func Test_palindromeNum(t *testing.T) {
	ast := assert.New(t)

	for _, tc :=range tcs {
		rescnt := palindromeNum(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
