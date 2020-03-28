package reverse_linked_list_206

import . "github.com/austingebauer/go-leetcode/structures"

// Returns a pointer to a list that is the reverse of the passed head.
// This function prepends the first n values in the passed list
// into a new list and returns the new list.
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var rev *ListNode
	for head != nil {
		// hold last head
		hold := head
		head = head.Next

		// prepend into rev
		if rev == nil {
			rev = hold
			rev.Next = nil
		} else {
			hold.Next = rev
			rev = hold
		}
	}

	return rev
}

// Reverses the passed list without creating a new list by
// prepending front nodes to the end.
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// find the end of the list
	end := head
	for end.Next != nil {
		end = end.Next
	}

	// prepend every front node to the end
	// 1,2,3
	// 2,3,1
	// 3,2,1
	for end != head {
		hold := head
		head = head.Next
		hold.Next = end.Next
		end.Next = hold
	}

	return head
}
