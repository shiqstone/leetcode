package code020

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct{
	para string
	ans  bool
}{
	{
		"()[]{}",
		true,
	},
	{
		"(]",
		false,
	},
	{
		"({[]})",
		true,
	},
	{
		"(){[({[]})]}",
		true,
	},
	{
		"((([[[{{{",
		false,
	},
	{
		"(())]]",
		false,
	},
}

func Test_validParentheses(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		res := isValidParenteses(tc.para)
		ast.Equal(tc.ans, res, "input:%v", tc)
	}
}