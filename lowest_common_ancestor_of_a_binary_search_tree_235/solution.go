package lowest_common_ancestor_of_a_binary_search_tree_235

import . "github.com/austingebauer/go-leetcode/structures"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != p && root != q {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			// LCA is root because p and q are on different
			// sides (one larger, one smaller) of root
			return root
		}
	}

	return root
}
