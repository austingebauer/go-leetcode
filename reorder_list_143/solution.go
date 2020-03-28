package reorder_list_143

import . "github.com/austingebauer/go-leetcode/structures"

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	cur := head
	var end *ListNode
	for end != cur {
		end = head

		// move e until one before the end
		for end.Next != nil && end.Next.Next != nil {
			end = end.Next
		}

		if end != cur {
			end.Next.Next = cur.Next
			cur.Next = end.Next
			end.Next = nil
			cur = cur.Next.Next
		}
	}
}
