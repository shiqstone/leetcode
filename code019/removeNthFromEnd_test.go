package code019

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"src/leetcode/utils"
)

type para struct {
	arr []int
	pos int
}
var tcs = []struct {
	para para
	ans  []int
}{
	{
		para{
			[]int{1, 2}, 1,
		},
		[]int{1},
	},
	{
		para{
			[]int{1, 2, 3, 4, 5}, 1,
		},
		[]int{1, 2, 3, 4},
	},
	{
		para{
			[]int{1, 2,3, 4, 5}, 2,
		},
		[]int{1, 2, 3, 5},
	},
	{
		para{
			[]int{1}, 1,
		},
		[]int{},
	},
}

func Test_RemoveNthFromEnd(t *testing.T){
	ast := assert.New(t)

	for _, tc := range tcs {
		lst := slice2list(tc.para.arr)
		res := removeNthFromEnd(lst, tc.para.pos)
		slices := list2slice(res)
		ast.Equal(tc.ans, slices, "input:%v", tc)
	}
}

func slice2list(input []int) *utils.ListNode {
	if len(input) == 0 {
		return nil
	}
	res := &utils.ListNode{Val:input[0]}
	temp := res
	for i:=1; i<len(input); i++ {
		temp.Next = &utils.ListNode{Val:input[i]}
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