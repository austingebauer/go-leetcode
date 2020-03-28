package compare_strings_by_frequency_of_the_smallest_character_1170

// Note: study again.
//
// A good problem for runtime optimization and memory trade-offs.
// By taking memory, we can speed up runtime of the algorithm.
func numSmallerByFrequency(queries []string, words []string) []int {
	ans := make([]int, 0)
	queriesFreq := make([]int, len(queries))
	wordsFreq := make([]int, len(words))

	// pre-compute the frequency of the smallest character in queries and words
	for i, q := range queries {
		queriesFreq[i] = frequencyOfSmallestChar(q)
	}
	for i, w := range words {
		wordsFreq[i] = frequencyOfSmallestChar(w)
	}

	for _, qf := range queriesFreq {
		total := 0
		for _, wf := range wordsFreq {
			if qf < wf {
				total++
			}
		}
		ans = append(ans, total)
	}

	return ans
}

func frequencyOfSmallestChar(s string) int {
	if s == "" {
		return 0
	}

	smallest := string(s[0])
	frequency := 1
	for i := 1; i < len(s); i++ {
		if string(s[i]) < smallest {
			smallest = string(s[i])
			frequency = 1
		} else if string(s[i]) == smallest {
			frequency++
		}
	}

	return frequency
}

/*
// Note: started here with computing frequency every iteration of the loop
func numSmallerByFrequency(queries []string, words []string) []int {
	ans := make([]int, 0)

	for _, q := range queries {
		total := 0
		// Note: second optimization went below to replace LHS of <
		fq := frequencyOfSmallestChar(q)
		for _, w := range words {
			if frequencyOfSmallestChar(q) < frequencyOfSmallestChar(w) {
				total++
			}
		}
		ans = append(ans, total)
	}

	return ans
}
*/
