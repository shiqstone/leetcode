package code003

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct{
	str string
	ans int
}{
	{
		"abcabcbb",
		3,
	},
	{
		"bbbbb",
		1,
	},
	{
		"pwwkew",
		3,
	},
	{
		"abcded",
		5,
	},
}

func Test_longestSubstring(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := LongestSubstring(tc.str)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
