package palindrome_linked_list_234

import (
	. "github.com/austingebauer/go-leetcode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * O(n^2) rt and O(1) space.
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	current := head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}

	current = head
	front := head
	for i := 0; i < length/2; i++ {
		for j := 0; j < (length-1)-(2*i); j++ {
			current = current.Next
		}

		if front.Val != current.Val {
			return false
		}

		front = front.Next
		current = front
	}

	return true
}
