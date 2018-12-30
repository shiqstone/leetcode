package code024

import "leetcode/utils"

func swapNodesInPairs(node *utils.ListNode) *utils.ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	newHead := node.Next
	node.Next = swapNodesInPairs(newHead.Next)
	newHead.Next = node

	return newHead
}