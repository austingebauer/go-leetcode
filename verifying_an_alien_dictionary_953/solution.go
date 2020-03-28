package verifying_an_alien_dictionary_953

func isAlienSorted(words []string, order string) bool {
	charToIndex := make(map[string]int)
	for idx, ch := range order {
		charToIndex[string(ch)] = idx
	}

	for i := 0; i < len(words)-1; i++ {
		if compareAlienStrings(words[i], words[i+1], charToIndex) == 1 {
			return false
		}
	}

	return true
}

// Returns -1 if s1 < s2, 0 if s1 == s2, and 1 if s1 > s2
func compareAlienStrings(s1, s2 string, charToIndex map[string]int) int {
	if s1 == s2 {
		return 0
	}

	i := 0
	for i < len(s1) && i < len(s2) {
		if charToIndex[string(s1[i])] > charToIndex[string(s2[i])] {
			return 1
		}

		if charToIndex[string(s1[i])] < charToIndex[string(s2[i])] {
			return -1
		}

		i++
	}

	// s2 is a prefix of s1
	if i < len(s1) {
		return 1
	}

	// s1 is a prefix of s2
	return -1
}
