package merge_k_sorted_lists_23

import . "github.com/austingebauer/go-leetcode/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// My first solution to divide and conquer.
func mergeKListsDivideAndConquer(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	k := len(lists)
	for k > 1 {
		acc := 0
		for i := 0; i < k; i += 2 {
			if i+1 == k {
				lists[i-acc] = lists[i]
				break
			}

			lists[i-acc] = mergeTwoLists(lists[i], lists[i+1])
			acc++
		}

		if k%2 == 1 {
			k = k/2 + 1
		} else {
			k /= 2
		}
	}

	return lists[0]
}

// Clever solution to divide and conquer found on leetcode.
// Debugged to understand more.
func mergeKListsDivideAndConquerClever(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	amount := len(lists)
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}

func mergeKLists(lists []*ListNode) *ListNode {
	var merged *ListNode

	for i := 0; i < len(lists); i++ {
		merged = mergeTwoLists(merged, lists[i])
	}

	return merged
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
