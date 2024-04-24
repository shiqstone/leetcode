package code086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// Dummy nodes for partitions
	less := &ListNode{}
	greaterOrEqual := &ListNode{}

	// Pointers for partitions
	lessPtr := less
	greaterOrEqualPtr := greaterOrEqual

	// Traverse the original linked list
	for head != nil {
		if head.Val < x {
			lessPtr.Next = head
			lessPtr = lessPtr.Next
		} else {
			greaterOrEqualPtr.Next = head
			greaterOrEqualPtr = greaterOrEqualPtr.Next
		}
		head = head.Next
	}

	// Connect the two partitions
	lessPtr.Next = greaterOrEqual.Next
	greaterOrEqualPtr.Next = nil

	return less.Next
}

func createLinkedList(arr []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range arr {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}
