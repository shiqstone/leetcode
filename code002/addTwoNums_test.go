package code002

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"newstart/src/leetcode/utils"
)

var tcs = []struct{
	l1 []int
	l2 []int
	ans []int
}{
	{
		[]int{7,2,4,3},
		[]int{4,5,6,4},
		[]int{1,1,8,0,7},
	},
	{
		[]int{7,2,4,3},
		[]int{1,5,6,4},
		[]int{8,8,0,7},
	},
	{
		[]int{2,4,3},
		[]int{5,6,4},
		[]int{8,0,7},
	},
}

func Test_addTwoNums(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("---%v~~~\n", tc)
		l1 := utils.Ints2List(tc.l1)
		l2 := utils.Ints2List(tc.l2)
		reslst := addTwoNums(l1, l2)
		ares := utils.List2Ints(reslst)
		ast.Equal(tc.ans, ares, "input:%v", tc)
	}
}
