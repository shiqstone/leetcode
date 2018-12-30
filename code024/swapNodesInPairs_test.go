package code024

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"leetcode/utils"
)

var tcs = []struct{
	para []int
	ans []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{2, 1, 4, 3},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 1, 4, 3, 5},
	},
}

func Test_swapNodesInPairs(t *testing.T)  {
	ast := assert.New(t)

	for _, tc := range tcs {
		lst := utils.Ints2List(tc.para)
		res := swapNodesInPairs(lst)
		slices := utils.List2Ints(res)
		ast.Equal(tc.ans, slices, "input:%v", tc)
	}
}