package code019

import "src/leetcode/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	newHead := &utils.ListNode{}
	newHead.Next = head
	point1 := newHead
	point2 := newHead
	for ; point1 != nil; n-- {
		if n < 0 {
			point2 = point2.Next
		}
		point1 = point1.Next
	}
	point2.Next = point2.Next.Next
	return newHead.Next
}
