package code028

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct{
	haystack string
	needle string
	ans int
}{
	{
		"hello",
		"ll",
		2,
	},
	{
		"aaaa",
		"bba",
		-1,
	},
	{
		"",
		"int",
		0,
	},
}

func Test_implementStrstr(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		res := implementStrstr(tc.haystack, tc.needle)
		ast.Equal(tc.ans, res, "Input: %V", tc)
	}
}