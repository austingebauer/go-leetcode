package distribute_candies_575

import "math"

func distributeCandies(candies []int) int {
	cm := make(map[int]int)
	for _, v := range candies {
		_, ok := cm[v]
		if ok {
			cm[v]++
		} else {
			cm[v] = 1
		}
	}

	// Key idea here is that the maximum number of candies to
	// distribute to the sister is the minimum between half of
	// all candies (even number of candies assumed) AND the number
	// of unique candies available. She can have at most unique number
	// of candies unless there are more unique number of candies than
	// half of them.
	return int(math.Min(float64(len(cm)), float64(len(candies)/2)))
}

/*
// Note: old code
func distributeCandies(candies []int) int {
	cm := make(map[int]int)
	for _, v := range candies {
		_, ok := cm[v]
		if ok {
			cm[v]++
		} else {
			cm[v] = 1
		}
	}

	sisterMap := make(map[int]bool)
	for i := 0; i < len(candies) / 2; i++ {
		distributed := false
		for k, v := range cm {
			// No more of that candy left
			if v == 0 {
				continue
			}

			// Give sister a candy
			cm[k] -= 1
			sisterMap[k] = true

			// If sister has taken half of all candies, she's done
			if len(sisterMap) == len(candies) / 2 {
				distributed = true
				break
			}
		}
		if distributed {
			break
		}
	}

	return len(sisterMap)
}
*/
