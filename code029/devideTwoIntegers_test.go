package code029

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

type para = struct{
	m int
	n int
};

var tcs = []struct{
	para para
	ans int
}{
	{
		para{2, 3},
		0,
	},
	{
		para{2, 0},
		math.MaxInt32,
	},
	{
		para{1024, 3},
		341,
	},
	{
		para{-1024, 3},
		-341,
	},
	{
		para{-1024, -3},
		341,
	},
}

func Test_DevideTwoIntegers(t *testing.T){
	ast := assert.New(t)
	for _, tc := range tcs {
		res := devide(tc.para.m, tc.para.n)
		ast.Equal(tc.ans, res, "Input: %V", tc)
	}
}