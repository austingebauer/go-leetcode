package first_unique_character_in_a_string_387

import "math"

func firstUniqChar(s string) int {
	charMap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		idx, ok := charMap[s[i]]

		if ok && idx != -1 {
			charMap[s[i]] = -1
		}

		if !ok {
			charMap[s[i]] = i
		}
	}

	allNonUnique := true
	minUnique := math.MaxInt32
	for _, val := range charMap {
		if val != -1 {
			allNonUnique = false
			minUnique = int(math.Min(float64(minUnique), float64(val)))
		}
	}

	if allNonUnique {
		return -1
	}

	return minUnique
}
