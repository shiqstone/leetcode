package code013

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para  string
	ans   int
}{
	{
		"III",
		3,
	},
	{
		"IV",
		4,
	},
	{
		"IX",
		9,
	},
	{
		"XIII",
		13,
	},
	{
		"XXIX",
		29,
	},
	{
		"LVIII",
		58,
	},
	{
		"MDCCCLXXXVIII",
		1888,
	},
	{
		"MCMXCIV",
		1994,
	},
	{
		"MMMCMXCIX",
		3999,
	},
}

func Test_romanToIn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := romanToInt(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}