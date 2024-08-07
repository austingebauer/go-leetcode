package maximum_average_subarray_i_643

import "math"

func findMaxAverage(nums []int, k int) float64 {
	start := 0
	end := k - 1
	sum := 0.0

	for i := start; i <= end; i++ {
		sum += float64(nums[i])
	}

	maxAvg := sum / float64(k)
	for end < len(nums)-1 {
		sum -= float64(nums[start])
		sum += float64(nums[end+1])
		maxAvg = math.Max(maxAvg, sum/float64(k))
		start++
		end++
	}

	return maxAvg
}
