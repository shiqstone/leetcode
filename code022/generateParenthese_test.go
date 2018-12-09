package code022

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para int
	ans  []string
}{
	{
		2,
		[]string{
			"(())",
			"()()",
		},
	},
	{
		3,
		[]string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		},
	},
	{
		4,
		[]string{
			"(((())))",
			"((()()))",
			"((())())",
			"((()))()",
			"(()(()))",
			"(()()())",
			"(()())()",
			"(())(())",
			"(())()()",
			"()((()))",
			"()(()())",
			"()(())()",
			"()()(())",
			"()()()()",
		},
	},
}

func Test_generate(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		res := generateParenthese(tc.para)
		ast.Equal(tc.ans, res, "input:%V", tc)
	}
}