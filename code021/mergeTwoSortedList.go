package code021

import "leetcode/utils"

func mergeSortedList(lst1, lst2 *utils.ListNode) *utils.ListNode {
	if lst1 == nil {
		return lst2
	}
	if lst2 == nil {
		return lst1
	}
	var head, node *utils.ListNode
	if lst1.Val < lst2.Val {
		head = lst1
		node = lst1
		lst1 = lst1.Next
	} else {
		head = lst2
		node = lst2
		lst2 = lst2.Next
	}

	for lst1 != nil && lst2 != nil {
		if lst1.Val < lst2.Val {
			node.Next = lst1
			lst1 = lst1.Next
		} else {
			node.Next = lst2
			lst2 = lst2.Next
		}
		node = node.Next
	}

	if lst1 != nil {
		node.Next = lst1
	}
	if lst2 != nil {
		node.Next = lst2
	}
	return head
}
