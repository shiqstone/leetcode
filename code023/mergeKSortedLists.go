package code023

import "leetcode/utils"

func mergeKSortedList(karr []*utils.ListNode) *utils.ListNode {
	length := len(karr)
	half := length/2

	if length == 0 {
		return nil
	}
	if length == 1 {
		return karr[0]
	}

	if length == 2 {
		var (
			lst0, lst1 = karr[0], karr[1]
			res, cur *utils.ListNode
		)

		if lst0 == nil {
			return lst1
		}
		if lst1 == nil {
			return lst0
		}

		if lst0.Val < lst1.Val {
			res, cur, lst0 = lst0, lst0, lst0.Next
		} else {
			res, cur, lst1 = lst1, lst1, lst1.Next
		}

		for lst0 != nil && lst1 != nil {
			if lst0.Val < lst1.Val {
				cur.Next = lst0
				lst0 = lst0.Next
			} else {
				cur.Next = lst1
				lst1 = lst1.Next
			}
			cur = cur.Next
		}

		if lst0 != nil {
			cur.Next = lst0
		}
		if lst1 != nil {
			cur.Next = lst1
		}
		return res
	}

	first := mergeKSortedList(karr[half:])
	last  := mergeKSortedList(karr[:half])
	return mergeKSortedList([]*utils.ListNode{first, last})
}
