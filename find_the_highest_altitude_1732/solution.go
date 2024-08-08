package find_the_highest_altitude_1732

import "math"

// -4 -3 -2 -1   4  3  2
//
//	^
//
// -4 -7 -9 -10 -6 -3  -1
func largestAltitude(gain []int) int {
	prefixSums := make([]int, len(gain))
	for i, g := range gain {
		toAdd := 0
		if i > 0 {
			toAdd = prefixSums[i-1]
		}
		prefixSums[i] = toAdd + g
	}

	// catch is that zero starting point could be the highest altitude
	highest := 0
	for _, s := range prefixSums {
		highest = int(math.Max(float64(highest), float64(s)))
	}
	return highest
}
