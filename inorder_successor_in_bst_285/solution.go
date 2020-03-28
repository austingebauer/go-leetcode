package inorder_successor_in_bst_285

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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var successor *TreeNode
	for root != nil {
		if root.Val > p.Val {
			// successor of a node p is the node with
			// the smallest key greater than p.val.
			if successor == nil || root.Val < successor.Val {
				successor = root
			}
		}

		if root.Val > p.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return successor
}
