package insert_interval_57

import (
	"math"
)

// Note: cool hard problem :)
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	merged := make([][]int, 0)

	// discard all non-overlapping intervals by placing them into merged
	i := 0
	for i < len(intervals) {
		if intervals[i][1] < newInterval[0] {
			merged = append(merged, intervals[i])
		} else {
			break
		}

		i++
	}

	// 0: did not find an interval where new interval start is less than i-th interval end
	if i == len(intervals) {
		return append(merged, newInterval)
	}

	// 1: interval doesn't merge at all
	if newInterval[0] < intervals[i][0] && newInterval[1] < intervals[i][0] {
		return append(append(merged, newInterval), intervals[i:]...)
	}

	// the minimum merged value is the new interval start or the i-th interval start
	minMerged := int(math.Min(float64(newInterval[0]), float64(intervals[i][0])))

	// 2: the new interval end is between the i-th start and end inclusive
	if newInterval[1] >= intervals[i][0] && newInterval[1] <= intervals[i][1] {
		maxMerged := intervals[i][1]
		merged = append(merged, []int{minMerged, maxMerged})
		i++
	} else {
		// 3: the new interval end could span many intervals
		maxMerged := -1
		i++
		for i < len(intervals) {
			if newInterval[1] < intervals[i][0] {
				maxMerged = newInterval[1]
				break
			} else if newInterval[1] >= intervals[i][0] && newInterval[1] <= intervals[i][1] {
				// if new interval end falls in between an interval
				maxMerged = intervals[i][1]
				i++
				break
			}

			i++
		}

		// new interval end is larger than all other intervals
		if maxMerged == -1 {
			maxMerged = newInterval[1]
		}

		merged = append(merged, []int{minMerged, maxMerged})
	}

	// place the tail (n + 1) intervals into merged
	for i < len(intervals) {
		merged = append(merged, intervals[i])
		i++
	}

	return merged
}

// Note: good attempt with 152 / 154 test cases passed.
// This got messy. Tried to merge newInterval into intervals. Can do better!
func insert0(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	// interval goes in front and doesn't merge at all
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}

	// interval goes in front and merges
	if newInterval[1] == intervals[0][0] {
		intervals[0][0] = newInterval[0]
		return intervals
	}

	// newInterval start is less than first interval start
	if newInterval[0] < intervals[0][0] {
		intervals[0][0] = newInterval[0]
		intervals[0][1] = int(math.Max(float64(newInterval[1]), float64(intervals[0][1])))
	}

	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]

		// find where the start is less than the finish of interval
		if newInterval[0] <= interval[1] {
			// while newIntervals end is greater than starts of subsequent intervals
			j := i
			for j < len(intervals) {
				// insert in between j and j-1
				if i > 0 && intervals[j-1][1] < newInterval[0] && intervals[j][0] > newInterval[1] {
					return append(append(append([][]int{}, intervals[:j]...), newInterval), intervals[j:]...)
				}

				if newInterval[1] <= intervals[j][0] {
					break
				}

				j++
			}

			if j == len(intervals) {
				intervals[i][0] = int(math.Min(float64(intervals[i][0]),
					float64(newInterval[0])))
				intervals[i][1] = int(math.Max(float64(newInterval[1]),
					float64(intervals[len(intervals)-1][1])))
				intervals = intervals[0 : i+1]
			} else if newInterval[1] < intervals[j][0] {
				intervals[i][0] = int(math.Min(float64(intervals[i][0]),
					float64(newInterval[0])))
				intervals[i][1] = int(math.Max(float64(intervals[j-1][1]),
					float64(newInterval[1])))
				if j-i > 1 {
					intervals = append(append([][]int{}, intervals[0:j-1]...), intervals[j:]...)
				}
			} else if newInterval[1] <= intervals[j][0] {
				intervals[i][1] = intervals[j][1]
				intervals = append(append([][]int{}, intervals[0:i+1]...), intervals[j+1:]...)
			}

			return intervals
		}
	}

	return append(intervals, newInterval)
}
