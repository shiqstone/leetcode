package code026

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct{
	para []int
	ans int
}{
	{
		[]int{1, 1, 2},
		2,
	},
	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},
}

func Test_removeDuplicatesFromSortedArray(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		res := removeDuplicates(tc.para)
		ast.Equal(tc.ans, res, "Input: %v", tc)
	}
}