package intersection_of_two_linked_lists_160

import (
	"testing"

	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	intersectNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectNode}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersectNode}}}
	assert.Equal(t, intersectNode, getIntersectionNode(headA, headB))
}
