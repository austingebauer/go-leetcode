package meeting_rooms_252

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	// Sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// If any intervals overlap, one person could not attend all meetings
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}
