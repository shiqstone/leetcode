package code010

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type para struct {
	input string
	parten string
}

var tcs = []struct {
	p para
	ans bool
} {
	{
		para {
			input:"a",
			parten:"aa",
		},
		false,
	},
	{
		para {
			"aa",
			"aa",
		},
		true,
	},
	{
		para {
			"aaa",
			"a*",
		},
		true,
	},
	{
		para {
			"ab",
			".*",
		},
		true,
	},
	{
		para {
			"aab",
			"c*a*b",
		},
		true,
	},
	{
		para {
			"mississippi",
			"mis*is*p*.",
		},
		false,
	},
}

func Test_regularMatch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		rescnt := isMatch(tc.p.input, tc.p.parten)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
