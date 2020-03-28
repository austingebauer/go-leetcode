package maximum_subarray_53

import "math"

// Note: study again
// Kadane's Algorithm works for this problem
// https://www.youtube.com/watch?v=86CQq3pKSUw
// Remembering past sums (dynamic programming) reduces runtime.
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	maxCurrent := nums[0]

	for i := 1; i < len(nums); i++ {
		maxCurrent = int(math.Max(float64(nums[i]), float64(maxCurrent+nums[i])))
		maxSum = int(math.Max(float64(maxSum), float64(maxCurrent)))
	}

	return maxSum
}

// Solved for the 2nd time
func maxSubArray2(nums []int) int {
	maxG := math.MinInt32
	maxL := math.MinInt32
	for i := 0; i < len(nums); i++ {
		maxL = int(math.Max(float64(nums[i]), float64(maxL+nums[i])))
		maxG = int(math.Max(float64(maxG), float64(maxL)))
	}

	return maxG
}
