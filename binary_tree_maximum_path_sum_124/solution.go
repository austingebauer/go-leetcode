package binary_tree_maximum_path_sum_124

import (
	. "github.com/austingebauer/go-leetcode/structures"
	"math"
)

var max int

func maxPathSum(root *TreeNode) int {
	max = math.MinInt32
	maxPathSumHelper(root)
	return max
}

func maxPathSumHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Including the left/right child in the current sum is not optimal
	// if either is negative, so take the max of the left/right sum and 0
	leftSum := int(math.Max(float64(maxPathSumHelper(root.Left)), 0.0))
	rightSum := int(math.Max(float64(maxPathSumHelper(root.Right)), 0.0))
	pathSum := rightSum + leftSum + root.Val
	max = int(math.Max(float64(max), float64(pathSum)))

	return root.Val + int(math.Max(float64(rightSum), float64(leftSum)))
}
