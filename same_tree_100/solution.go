package same_tree_100

import . "github.com/austingebauer/go-leetcode/structures"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// if p is nil and q is not or vice versa
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// p and q are either both nil or both not nil
	if p == nil && q == nil {
		return true
	}

	// p and q are both not nil
	return p.Val == q.Val &&
		isSameTree(p.Right, q.Right) &&
		isSameTree(p.Left, q.Left)
}
