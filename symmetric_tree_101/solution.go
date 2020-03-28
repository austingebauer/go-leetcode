package symmetric_tree_101

import . "github.com/austingebauer/go-leetcode/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Note: good problem! Review this.
// Use of two pointers to left and right subtree was a new idea to me.
// This allows us to check each side of the tree, which is different to
// other recursive (pre, in, post) traversals that I'm familiar with.
func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	// if left XOR right is nil (check above makes XOR)
	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) &&
		isSym(right.Left, left.Right) &&
		isSym(left.Left, right.Right)
}

// Based on in-order traversal being a palindrome.
// Wrong approach but good lesson on array appending in recursive function.
func isSymmetricWrong(root *TreeNode) bool {
	inOrder := make([]int, 0)
	inOrder = inorderTraverse(root, inOrder)

	i := 0
	j := len(inOrder) - 1
	for i < j {
		if inOrder[i] != inOrder[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func inorderTraverse(root *TreeNode, a []int) []int {
	if root == nil {
		return a
	}

	a = append(inorderTraverse(root.Left, a), root.Val)
	return inorderTraverse(root.Right, a[:])
}
