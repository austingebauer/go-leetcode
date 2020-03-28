package maximum_product_subarray_152

import "math"

// Note: study again.
// Time: O(n), Space: O(1)
func maxProduct(nums []int) int {
	minCurrent := nums[0]
	maxCurrent := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		// if the next number is negative, then swap minCurrent and maxCurrent
		// so that we can maximize the swapping signs of our most minimum number
		if nums[i] < 0 {
			minCurrent, maxCurrent = maxCurrent, minCurrent
		}

		// similar to Kadane's algorithm, with small twist for negative number products
		minCurrent = int(math.Min(float64(nums[i]), float64(minCurrent*nums[i])))
		maxCurrent = int(math.Max(float64(nums[i]), float64(maxCurrent*nums[i])))
		max = int(math.Max(float64(max), float64(maxCurrent)))
	}

	return max
}

// First, O(n^2) solution.
func maxProduct0(nums []int) int {
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		product := nums[i]
		max = int(math.Max(float64(max), float64(product)))
		for j := i + 1; j < len(nums); j++ {
			product *= nums[j]
			max = int(math.Max(float64(max), float64(product)))
		}
	}

	return max
}
