package code011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para  []int
	ans   int
}{
	{
		[]int{1,8,6,2,5,4,8,3,7},
		49,
	},
	{
		[]int{1, 3, 6, 4, 3, 5, 6, 7, 8, 9, 7, 5, 4, 3, 2, 1},
		48,
	},
	{
		[]int{1, 2, 3, 1},
		3,
	},
}

func Test_reverseInt(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := maxArea(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
