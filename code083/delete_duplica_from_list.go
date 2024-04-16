package code083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		// Check if current has duplicates
		if current.Val == current.Next.Val {
			// Remove duplicate node by updating links
			current.Next = current.Next.Next
		} else {
			// Move to the next node
			current = current.Next
		}
	}

	return head
}

// Function to create a linked list from an array
func createLinkedList(arr []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range arr {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}
