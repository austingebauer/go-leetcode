package template

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
	max := 0
	for _, s := range prefixSums {
		max = int(math.Max(float64(max), float64(s)))
	}
	return max
}
