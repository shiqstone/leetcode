package code021

import (
	"github.com/stretchr/testify/assert"
	"leetcode/utils"
	"testing"
)

type para struct {
	one []int
	two []int
}

var tcs = []struct {
	para para
	ans  []int
}{
	{
		para{
			[]int{1, 3, 5, 7},
			[]int{},
		},
		[]int{1, 3, 5, 7},
	},
	{
		para{
			[]int{1, 3, 5, 7},
			[]int{2, 4, 6, 8},
		},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		para{
			[]int{10, 20, 30},
			[]int{1, 2, 3},
		},
		[]int{1, 2, 3, 10, 20, 30},
	},
	{
		para{
			[]int{1, 3, 5},
			[]int{2, 4, 6, 8, 10},
		},
		[]int{1, 2, 3, 4, 5, 6, 8, 10},
	},
	{
		para{
			[]int{1, 3, 5, 7, 9},
			[]int{2, 4, 6},
		},
		[]int{1, 2, 3, 4, 5, 6, 7, 9},
	},
}

func Test_mergeList(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		lst1 := slice2list(tc.para.one)
		lst2 := slice2list(tc.para.two)

		res := mergeSortedList(lst1, lst2)
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
