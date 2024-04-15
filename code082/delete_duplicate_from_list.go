package code082

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	current := head

	for current != nil {
		// Check if current has duplicates
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}

		// If current node is unique, update links
		if prev.Next == current {
			prev = current
		} else {
			prev.Next = current.Next
		}

		current = current.Next
	}

	return dummy.Next
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

// Function to print a linked list
// func printLinkedList(head *ListNode) {
// 	current := head
// 	for current != nil {
// 		fmt.Print(current.Val, " -> ")
// 		current = current.Next
// 	}
// 	fmt.Println("nil")
// }
