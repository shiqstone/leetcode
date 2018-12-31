package code025

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"leetcode/utils"
)

var tcs = []struct{
	para []int
	k int
	ans []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		3,
		[]int{3, 2, 1, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		2,
		[]int{2, 1, 4, 3, 5},
	},
}

func Test_reverseNodesInKGroup(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		lst := utils.Ints2List(tc.para)
		res := reverseKGroup(lst, tc.k)
		slices := utils.List2Ints(res)
		ast.Equal(tc.ans, slices, "Input: %v", tc)
	}
}
