package validate_binary_search_tree_98

import . "github.com/austingebauer/go-leetcode/structures"

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	in := inorderTraversal(root)
	for i := 1; i < len(in); i++ {
		if in[i-1] >= in[i] {
			return false
		}
	}

	return true
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	o := inorderTraversal(root.Left)
	o = append(o, root.Val)
	o = append(o, inorderTraversal(root.Right)...)
	return o
}
