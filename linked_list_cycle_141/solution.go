package linked_list_cycle_141

import . "github.com/austingebauer/go-leetcode/structures"

// Time: O(n)
// Space: O(1)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// use the rabbit and turtle analogy
	// if there as a cycle, eventually the rabbit will meet back up with the turtle
	turtle := head
	rabbit := head.Next
	for rabbit != nil && turtle != nil {
		if turtle == rabbit {
			return true
		}

		if rabbit.Next == nil {
			rabbit = nil
		} else {
			rabbit = rabbit.Next.Next
		}

		if turtle.Next == nil {
			turtle = nil
		} else {
			turtle = turtle.Next
		}
	}

	return false
}
