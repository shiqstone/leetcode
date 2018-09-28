package code004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums1 []int
	nums2 []int
	ans   float64
}{
	{
		[]int{1, 3},
		[]int{2},
		2.0,
	},
	{
		[]int{1, 2},
		[]int{3, 4},
		2.5,
	},
	{
		[]int{2, 4, 6, 8},
		[]int{6, 9},
		6.0,
	},
}

func Test_addTwoNums(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := medianOfTwoSortedArray(tc.nums1, tc.nums2)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
