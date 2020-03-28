package subtree_of_another_tree_572

import (
	. "github.com/austingebauer/go-leetcode/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	return equalTrees(s, t) ||
		isSubtree(s.Left, t) ||
		isSubtree(s.Right, t)
}

func equalTrees(t1, t2 *TreeNode) bool {
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}

	if t1 == nil && t2 == nil {
		return true
	}

	return t1.Val == t2.Val &&
		equalTrees(t1.Left, t2.Left) &&
		equalTrees(t1.Right, t2.Right)
}
