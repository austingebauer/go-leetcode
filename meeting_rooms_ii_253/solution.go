package meeting_rooms_ii_253

import (
	"sort"
)

// First time around was more elegant, although, I didn't solve it without help
// the first time. I solved this again, reasoning alone, and came up with solution
// below. The key idea here is to understand that we are trying to slot the soonest
// starting meeting into the same room as the soonest ending meeting. If we can't
// (overlaps), then we need a new room.
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	starts := intervals
	ends := make([][]int, len(intervals))
	for i, v := range intervals {
		ends[i] = v
	}

	// sort the intervals by starting time to produce
	// ordering of meetings starting the soonest
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})

	// sort the intervals by ending time to produce
	// ordering of meetings ending the soonest
	sort.Slice(ends, func(i, j int) bool {
		return ends[i][1] < ends[j][1]
	})

	// try to book the soonest starting meetings into
	// the soonest ending rooms. If we can't, then we
	// need another meeting room.
	rooms := 0
	end := 0
	start := 0
	for start < len(starts) && end < len(ends) {
		if starts[start][0] == ends[end][0] && starts[start][1] == ends[end][1] {
			// if it's the same meeting
			rooms++
		} else if starts[start][0] < ends[end][1] {
			// if soonest starting meeting overlaps with soonest ending,
			// then we need a new room
			rooms++
		} else if starts[start][0] >= ends[end][1] {
			// if soonest starting meeting doesn't overlap with soonest ending,
			// we can reuse the soonest ending room
			end++
		}

		start++
	}

	return rooms
}

// Note: Study again. Really good problem!
func minMeetingRooms0(intervals [][]int) int {
	sTimes := make([]int, len(intervals))
	eTimes := make([]int, len(intervals))
	for idx, v := range intervals {
		sTimes[idx] = v[0]
		eTimes[idx] = v[1]
	}

	sort.Slice(sTimes, func(i, j int) bool {
		return sTimes[i] < sTimes[j]
	})
	sort.Slice(eTimes, func(i, j int) bool {
		return eTimes[i] < eTimes[j]
	})

	var sIdx, eIdx, rooms int
	for sIdx < len(sTimes) {
		// if the soonest meeting start time is before the
		// soonest meeting end time, we need a new room
		if sTimes[sIdx] < eTimes[eIdx] {
			rooms++
		} else {
			// we can reuse an existing room, so
			// go to the next soonest meeting ending
			eIdx++
		}

		sIdx++
	}

	return rooms
}
