package intersection_of_two_linked_lists_160

import (
	"math"

	. "github.com/austingebauer/go-leetcode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := length(headA), length(headB)

	// Set itr1 to the list with the greater length
	var itr1, itr2 *ListNode
	if lenA > lenB {
		itr1, itr2 = headA, headB
	} else {
		itr1, itr2 = headB, headA
	}

	// Advance itr1 so that it's in step with itr2 with respect to length
	distance := int(math.Abs(float64(lenA - lenB)))
	for i := 0; i < distance; i++ {
		itr1 = itr1.Next
	}

	// Move both pointers forward until an intersection occurs or end of list
	for itr1 != nil {
		if itr1 == itr2 {
			return itr1
		}
		itr1, itr2 = itr1.Next, itr2.Next
	}

	return nil
}

func length(n *ListNode) int {
	var l int
	for n != nil {
		l++
		n = n.Next
	}
	return l
}
