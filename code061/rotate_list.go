package code061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	// Calculate the length of the linked list and find the tail node
	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	// Connect the tail to the head to form a cycle
	tail.Next = head

	// Calculate the new head position after rotation
	newTailPosition := length - k%length

	// Traverse to the new tail position
	newTail := head
	for i := 1; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	// Set the next node after the new tail as the new head
	newHead := newTail.Next

	// Break the cycle
	newTail.Next = nil

	return newHead
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
