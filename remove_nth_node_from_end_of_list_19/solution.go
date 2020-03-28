package remove_nth_node_from_end_of_list_19

import . "github.com/austingebauer/go-leetcode/structures"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	n = df(head, n)

	// if the head needs to be removed, then 0 is returned
	if n == 0 {
		head = head.Next
	}

	return head
}

func df(curr *ListNode, n int) int {
	if curr == nil {
		return n
	}

	// Depth-first, where we subtract from n every time
	// time a stack frame is popped off.
	n = df(curr.Next, n)

	// If n is zero, we can skip over the curr.Next.Next,
	// which removes the nth from the bottom of the list
	if n == 0 {
		curr.Next = curr.Next.Next
	}

	return n - 1
}
