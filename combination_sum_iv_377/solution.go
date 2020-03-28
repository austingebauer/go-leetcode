package combination_sum_iv_377

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	// for each number up to and including target
	for i := 1; i <= target; i++ {

		// for each number in nums we can pick from
		for _, n := range nums {
			if n > i {
				continue
			}

			// dp[i] is equal to the sum of past
			// dp values by choosing each number
			// n to try to make up the current
			// target i. See notes for table
			// explanation.
			dp[i] += dp[i-n]
		}
	}

	return dp[len(dp)-1]
}
