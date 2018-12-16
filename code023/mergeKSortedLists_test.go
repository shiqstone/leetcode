package code023

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"leetcode/utils"
)

var tcs = []struct{
	para [][]int
	ans  []int
}{
	{
		[][]int{
			[]int{1, 4, 7},
			[]int{2, 5, 8},
			[]int{3, 6, 9},
		},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		[][]int{
			[]int{1, 4, 7},
			[]int{2, 5, 8},
			[]int{},
		},
		[]int{1, 2, 4, 5, 7, 8},
	},
	{
		[][]int{
			[]int{1, 4, 7},
			[]int{},
			[]int{2, 5, 8},
		},
		[]int{1, 2, 4, 5, 7, 8},
	},
	{
		[][]int{},
		[]int{},
	},
}

func Test_mergeList(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		var karr []*utils.ListNode
		for _, arr := range tc.para{
			lst := slice2list(arr)
			karr = append(karr, lst)
		}

		res := mergeKSortedList(karr)
		slices := list2slice(res)
		ast.Equal(tc.ans, slices, "input:%V", tc)
	}
}

func slice2list(input []int) *utils.ListNode {
	if len(input) == 0 {
		return nil
	}
	res := &utils.ListNode{Val: input[0]}
	temp := res
	for i := 1; i < len(input); i++ {
		temp.Next = &utils.ListNode{Val: input[i]}
		temp = temp.Next
	}
	return res
}

func list2slice(head *utils.ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}