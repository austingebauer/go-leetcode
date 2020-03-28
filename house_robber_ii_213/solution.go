package house_robber_ii_213

import "math"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// The max is the max between robbing houses
	// where the beginning and end do not connect
	// into the ring of houses. So, that means house
	// 0 to one from the end, and house 1 to the end.
	return int(math.Max(
		float64(robInRange(nums[0:len(nums)-1])),
		float64(robInRange(nums[1:]))))
}

func robInRange(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = int(math.Max(float64(dp[i]), float64(nums[i]+dp[i-1])))
	}

	return dp[len(dp)-1]
}
