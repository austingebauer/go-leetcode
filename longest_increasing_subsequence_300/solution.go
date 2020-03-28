package longest_increasing_subsequence_300

import "math"

// Note: study again. Good dp problem.
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] is equal to the length of the longest increasing
	// subsequence possible considering elements from [0,i)
	dp := make([]int, len(nums))
	dp[0] = 1

	maxResult := 1
	for i := 0; i < len(nums); i++ {

		// dp[i] = max(dp[j]) + 1, for j in [0,i), where nums[i] > nums[j]
		jMax := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				jMax = int(math.Max(float64(jMax), float64(dp[j])))
			}
		}
		dp[i] = jMax + 1

		// set the final max result as we see new jMax
		maxResult = int(math.Max(float64(maxResult), float64(dp[i])))
	}

	return maxResult
}

func lengthOfLISOther(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// return explore1(nums)

	// return explore2(nums, math.MinInt32, 0)

	// Since prevIdx starts at -1, we need an extra row in memo
	// Each row represents the previous max value found from previous
	// to the current, which is the column value.
	memo := make([][]int, len(nums)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(nums))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	return explore3(nums, -1, 0, memo)
}

// This approach is better (O(n^2)) runtime.
// It is accepted by leetcode.
// It's not exactly simple though and the intuition for the
// memoization is a bit confusing.
func explore3(nums []int, prevIdx int, currentIdx int, memo [][]int) int {
	if currentIdx == len(nums) {
		return 0
	}

	if memo[prevIdx+1][currentIdx] >= 0 {
		return memo[prevIdx+1][currentIdx]
	}

	max1 := 0
	if prevIdx < 0 || nums[prevIdx] < nums[currentIdx] {
		max1 = 1 + explore3(nums, currentIdx, currentIdx+1, memo)
	}

	max2 := explore3(nums, prevIdx, currentIdx+1, memo)
	memo[prevIdx+1][currentIdx] = int(math.Max(float64(max1), float64(max2)))
	return memo[prevIdx+1][currentIdx]
}

// This approach is also slow (O(2^n) runtime.
// It works, but times out on leetcode.
// We can do better.
func explore2(nums []int, previousVal int, currentIdx int) int {
	if currentIdx == len(nums) {
		return 0
	}

	max1 := 0
	if previousVal < nums[currentIdx] {
		max1 = 1 + explore2(nums, nums[currentIdx], currentIdx+1)
	}
	max2 := explore2(nums, previousVal, currentIdx+1)

	return int(math.Max(float64(max1), float64(max2)))
}

// This approach is slow (exponential runtime).
// It works, but times out on leetcode.
// We can do better.
func explore1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if isIncreasing(nums) {
		return len(nums)
	}

	maxLen := 0
	for i := range nums {
		newNums := append(append([]int{}, nums[0:i]...), nums[i+1:]...)
		if isIncreasing(newNums) {
			return len(newNums)
		}

		max := explore1(newNums)
		if max > maxLen {
			maxLen = max
		}
	}

	return maxLen
}

func isIncreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}

	return true
}
