package merge_two_sorted_lists_21

import . "github.com/austingebauer/go-leetcode/structures"

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	// Note: Cool technique to not have to worry about front of list being null
	//       current.Next is always where the next will go
	holdFront := &ListNode{}
	current := holdFront

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}

	return holdFront.Next
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	var merged *ListNode
	var mergedEnd *ListNode
	for l1 != nil && l2 != nil {
		var next *ListNode
		if l1.Val <= l2.Val {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}

		if merged == nil {
			merged = next
			mergedEnd = next
		} else {
			mergedEnd.Next = next
			mergedEnd = next
		}
	}

	if merged == nil {
		if l1 == nil {
			return l2
		}

		return l1
	}

	if l1 == nil && l2 != nil {
		mergedEnd.Next = l2
	}

	if l2 == nil && l1 != nil {
		mergedEnd.Next = l1
	}

	return merged
}

func mergeTwoLists0(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	merged := &ListNode{}
	current := merged

	for l1 != nil || l2 != nil {
		if l1 == nil {
			for l2 != nil {
				current.Val = l2.Val
				if l2.Next != nil {
					current.Next = &ListNode{}
					current = current.Next
				}
				l2 = l2.Next
			}

			return merged
		}
		if l2 == nil {
			for l1 != nil {
				current.Val = l1.Val
				if l1.Next != nil {
					current.Next = &ListNode{}
					current = current.Next
				}
				l1 = l1.Next
			}

			return merged
		}

		if l1.Val <= l2.Val {
			current.Val = l1.Val
			current.Next = &ListNode{}
			l1 = l1.Next
			current = current.Next
		} else {
			current.Val = l2.Val
			current.Next = &ListNode{}
			l2 = l2.Next
			current = current.Next
		}
	}

	return merged
}
