package kth_smallest_element_in_a_bst_230

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"math"
)

func kthSmallest(root *TreeNode, k int) int {
	_, v := inOrderTraverse(root, k)
	return v
}

func inOrderTraverse(root *TreeNode, k int) (int, int) {
	if root == nil {
		return k, math.MaxInt32
	}

	var val int
	k, val = inOrderTraverse(root.Left, k)

	// Stop traverse when the kth smallest number has been found
	if val != math.MaxInt32 {
		return k, val
	}

	// When k becomes zero, the kth smallest number has been found
	k--
	if k == 0 {
		return k, root.Val
	}

	return inOrderTraverse(root.Right, k)
}
