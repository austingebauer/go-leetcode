package maximum_depth_of_binary_tree_104

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"math"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(
		float64(1+maxDepth(root.Left)),
		float64(1+maxDepth(root.Right))))
}
