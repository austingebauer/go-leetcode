package linked_list_cycle_141

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasCycle1(t *testing.T) {
	four := &ListNode{
		Val:  -4,
		Next: nil,
	}
	two := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: four,
		},
	}
	four.Next = two

	head := &ListNode{
		Val:  3,
		Next: two,
	}

	assert.Equal(t, true, hasCycle(head))
}

func Test_hasCycle2(t *testing.T) {
	one := &ListNode{
		Val:  1,
		Next: nil,
	}
	two := &ListNode{
		Val:  2,
		Next: one,
	}
	one.Next = two

	assert.Equal(t, true, hasCycle(one))
}

func Test_hasCycl32(t *testing.T) {
	one := &ListNode{
		Val:  1,
		Next: nil,
	}

	assert.Equal(t, false, hasCycle(one))
}
