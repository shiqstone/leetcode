package code008

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str   string
	ans   int32
}{
	{
		"42",
		42,
	},
	{
		"   -42",
		-42,
	},
	{
		"4193 with words",
		4193,
	},
	{
		"words and 987",
		0,
	},
	{
		"-91283472332",
		-2147483648,
	},
}

func Test_StringToInteger(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := myAtoi(tc.str)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
