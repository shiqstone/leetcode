package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Ints(head *ListNode) []int {
	limit := 100
	times := 9

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("over link depth more than %d", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}

	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}
	return res
}
