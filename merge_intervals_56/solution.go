package merge_intervals_56

import (
	"math"
	"sort"
)

// Second solution
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	// sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	j := 0
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		prev := merged[j]
		next := intervals[i]

		// should we merge the interval
		if next[0] <= prev[1] {
			prev[1] = int(math.Max(float64(prev[1]), float64(next[1])))
		} else {
			merged = append(merged, next)
			j++
		}
	}

	return merged
}

// First solution
func merge0(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(merged) == 0 || intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
		} else if intervals[i][1] > merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = intervals[i][1]
		}
	}

	return merged
}
