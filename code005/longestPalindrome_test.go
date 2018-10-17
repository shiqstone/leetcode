package code005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str   string
	ans   string
}{
	{
		"babad",
		"bab",
	},
	{
		"cbbd",
		"bb",
	},
	{
		"abccba",
		"abccba",
	},
}

func Test_longestPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := longestPalindrome(tc.str)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
