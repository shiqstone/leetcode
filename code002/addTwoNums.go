package code002

import "newstart/src/leetcode/utils"

type ListNode = utils.ListNode

func addTwoNums(list1 *ListNode, list2 *ListNode) *ListNode {
	s1 := make([]int, 0, 32)
	for list1 !=nil {
		s1 = append(s1, list1.Val)
		list1 = list1.Next
	}

	s2 := make([]int, 0, 32)
	for list2 !=nil {
		s2 = append(s2, list2.Val)
		list2 = list2.Next
	}

	sum := 0
	head := &ListNode{Val: 0}

	for len(s1)>0 || len(s2)>0 {
		ls1 := len(s1)
		if ls1 >0 {
			sum += s1[ls1-1]
			s1 = s1[:ls1-1]
		}
		ls2 := len(s2)
		if ls2 >0 {
			sum += s2[ls2-1]
			s2 = s2[:ls2-1]
		}

		head.Val = sum % 10
		ln := &ListNode{Val : sum/10}
		ln.Next = head
		head = ln

		sum /= 10
	}

	if head.Val == 0 {
		return head.Next
	}
	return head
}
