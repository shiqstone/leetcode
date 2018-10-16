package code006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str   string
	rows  int
	ans   string
}{
	{
		"PAYPALISHIRING",
		3,
		"PAHNAPLSIIGYIR",
	},
	{
		"PAYPALISHIRING",
		4,
		"PINALSIGYAHRPI",
	},
	{
		"ABCDEFGHIJKLMNOPQRSTUVWX",
		5,
		"AIQBHJPRXCGKOSWDFLNTVEMU",
	},
	{
		"ABCDEFGHIJKLMN",
		4,
		"AGMBFHLNCEIKDJ",
	},
}

func Test_addTwoNums(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := convert(tc.str, tc.rows)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}


func Test_addTwoNums2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := convert2(tc.str, tc.rows)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
