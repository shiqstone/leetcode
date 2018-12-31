package code025

import "leetcode/utils"

func reverseKGroup(node *utils.ListNode, k int) *utils.ListNode {
	if k < 2 || node == nil || node.Next == nil {
		return node
	}

	tail, needReverse := getTail(node, k)

	if needReverse {
		tailNext := tail.Next
		// 斩断 tail 后的链接
		tail.Next = nil
		node, tail = revers(node, tail)
		// tail 后面接上尾部的递归处理
		tail.Next =reverseKGroup(tailNext, k)
	}
	return node
}

func getTail(node *utils.ListNode, k int) (*utils.ListNode, bool) {
	for k > 1 && node != nil {
		node = node.Next
		k--
	}
	return node, k==1 && node != nil
}

func revers(node, tail *utils.ListNode) (*utils.ListNode, *utils.ListNode) {
	curPre, cur := node, node.Next
	for cur != nil {
		curPre, cur, cur.Next = cur, cur.Next, curPre
	}
	return tail, node
}
