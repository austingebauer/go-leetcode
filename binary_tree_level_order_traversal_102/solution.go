package binary_tree_level_order_traversal_102

import . "github.com/austingebauer/go-leetcode/structures"

// Slight variation of first correct solution below, which
// doesn't keep counters for the number of elements in a
// level. It just uses the current, constant size of the queue
// (before anything new is added to it).
func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		// dequeue and create level for the length
		// of the currently enqueued level
		level := make([]int, 0)
		levelLen := len(q)
		for i := 0; i < levelLen; i++ {
			// add nodes value to the level
			n := q[0]
			level = append(level, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}

			// dequeue the node
			q = q[1:]
		}

		levels = append(levels, level)
	}

	return levels
}

// Note: study again.
func levelOrder0(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}

	levelCnt := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		// dequeue and create level levelCnt times
		newLevelCnt := 0
		level := make([]int, 0)
		for i := 0; i < levelCnt; i++ {
			// add nodes value to the level
			n := q[0]
			level = append(level, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
				newLevelCnt++
			}

			if n.Right != nil {
				q = append(q, n.Right)
				newLevelCnt++
			}

			// dequeue the node
			q = q[1:]
		}

		levels = append(levels, level)
		levelCnt = newLevelCnt
	}

	return levels
}
