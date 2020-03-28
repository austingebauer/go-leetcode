package letter_combinations_of_a_phone_number_17

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapping := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	combos := mapping[string(digits[0])]
	if len(digits) == 1 {
		return combos
	}

	// clip off the first digit since it's already in combos
	digits = digits[1:]
	for _, c := range digits {
		nextCombos := make([]string, 0)
		for j := 0; j < len(combos); j++ {
			for k := 0; k < len(mapping[string(c)]); k++ {
				nextCombos = append(nextCombos, combos[j]+mapping[string(c)][k])
			}
		}

		combos = nextCombos
	}

	return combos
}
