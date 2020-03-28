package non_overlapping_intervals_435

import (
	"math"
	"sort"
)

// DP solution based on sorting by start times.
// Almost identical solution to longest_increasing_subsequence_300, where
// dp[i] stores the max (longest) we *can keep* up to the i-th interval.
//
// Time: O(n^2), Space: O(n)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}

	// sort by interval start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// dp[i] stores the max amount of intervals we *can keep*
	// up to the i-th interval. dp[i + 1] is the (max amount from
	// dp[0, i]) + 1 which does not overlap with dp[i + 1].
	dp := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				max = int(math.Max(float64(max), float64(dp[j])))
			}
		}

		dp[i] = max + 1
	}

	// dp[len(dp)-1] holds the max intervals we can keep, so
	// the difference of all intervals and max we can keep is
	// the min we can erase
	return len(intervals) - dp[len(dp)-1]
}

// Second solution, which is correct but time limit exceeded.
func eraseOverlapIntervals1(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}

	// sort by interval start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return eraseOverlapHelper(intervals, 0)
}

func eraseOverlapHelper(intervals [][]int, count int) int {
	min := math.MaxInt32
	if !overlaps(intervals) {
		min = count
	}

	for i := 0; i < len(intervals); i++ {
		front := intervals[:i]
		back := intervals[i+1:]
		next := append(append([][]int{}, front...), back...)
		min = int(math.Min(float64(min), float64(eraseOverlapHelper(next, count+1))))
	}

	return min
}

func overlaps(intervals [][]int) bool {
	for j := 1; j < len(intervals); j++ {
		if intervals[j-1][1] > intervals[j][0] {
			return true
		}
	}

	return false
}

// Initial solution, which failed to recognize that the removal
// order isn't necessarily a part of the start sorted sequence.
func eraseOverlapIntervals0(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}

	// sort by interval start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	erase := 0
	overlapIdx := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[overlapIdx][1] > intervals[j][0] {
			erase++
		} else {
			overlapIdx++
		}
	}

	return erase
}
