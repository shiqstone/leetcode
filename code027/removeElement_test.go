package code027

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct{
	para []int
	k int
	ans []int
}{
	{
		[]int{3, 2, 2, 3},
		3,
		[]int{2, 2},
	},
	{
		[]int{3, 1, 5, 7, 2, 2, 3},
		3,
		[]int{2, 1, 5, 7, 2},
	},
	{
		[]int{1, 5, 7, 2, 2},
		3,
		[]int{1, 5, 7, 2, 2},
	},
}

func Test_removeElement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		res := removeElement(tc.para, tc.k)
		ast.Equal(tc.ans, res, "Input: %V", tc)
	}
}
