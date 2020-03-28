package construct_binary_tree_from_preorder_and_inorder_traversal_105

import . "github.com/austingebauer/go-leetcode/structures"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	// middle splits the inorder array into left and right subtrees of root
	middle := findIndex(inorder, root.Val)
	root.Left = buildTree(preorder[1:middle+1], inorder[:middle])
	root.Right = buildTree(preorder[middle+1:], inorder[middle+1:])
	return root
}

func findIndex(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}

	return -1
}

// First attempt, getting closer, but still didn't see key to unlock solving the problem.
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	return bt(preorder, inorder)
}

func bt(preorder []int, inorder []int) *TreeNode {
	if preorder[0] == inorder[0] {
		return &TreeNode{Val: preorder[0]}
	}

	curr := &TreeNode{}
	preorder = preorder[1:]
	curr.Left = bt(preorder, inorder)

	curr.Val = preorder[0]
	preorder = preorder[2:]
	inorder = inorder[2:]

	curr.Right = bt(preorder, inorder)

	return curr
}

// Initial solution, which was off. Mostly prototyping around incomplete idea.
func buildTree0(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	curr := root
	p := 1
	i := 0
	left := true
	for p < len(preorder) {
		if left {
			curr.Left = &TreeNode{Val: preorder[p]}
			curr = curr.Left
		} else {
			curr.Right = &TreeNode{Val: preorder[p]}
			curr = curr.Right
		}

		if preorder[p] != inorder[i] {
			p++
		} else {
			left = !left
			i += 2
		}
	}

	return root
}
