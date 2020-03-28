package minimum_window_substring_76

// Second solution. Intuition is to open the window
// until character frequencies in t are all less than zero.
// After, close the window until any single character frequency
// of t is larger than 0.
func minWindow(s string, t string) string {
	// create a map from character it's frequency in t
	tmap := make(map[string]int)
	for _, r := range t {
		tmap[string(r)] += 1
	}

	var min string
	var start, end int
	for start < len(s) && end < len(s) {
		_, ok := tmap[string(s[end])]
		if ok {
			tmap[string(s[end])]--

			// shrink start of window until window does not contain
			// every character in t
			for mapValuesLessThanOne(tmap) {
				// set the new min if it's smaller than old min
				if min == "" || len(min) > len(s[start:end+1]) {
					min = s[start : end+1]
				}

				// if s[start] is in tmap, increment it's count for removal
				_, ok := tmap[string(s[start])]
				if ok {
					tmap[string(s[start])]++
				}

				// shrink the start of the window
				start++
			}
		}

		end++
	}

	return min
}

func mapValuesLessThanOne(m map[string]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

// Initial solution. Pretty close, but overly complicated
// and assumed that duplicates of t were not allowed, meaning
// the substring of s must have exactly (and no more) the same
// frequency of unique characters in t.
func minWindow0(s string, t string) string {
	tmap := make(map[string]int)
	for _, r := range t {
		tmap[string(r)] += 1
	}

	// move i forward until a character in tmap is seen
	var i, j int
	for i = 0; i < len(s); i++ {
		_, ok := tmap[string(s[i])]
		if ok {
			break
		}
	}
	j = i

	min := ""
	for i < len(s) && j < len(s) {
		_, ok := tmap[string(s[j])]
		if ok {
			tmap[string(s[j])]--

			if tmap[string(s[j])] < 0 {
				for tmap[string(s[j])] < 0 {
					_, ok := tmap[string(s[i])]
					if ok {
						tmap[string(s[i])]++
					}
					i++
				}
				_, ok := tmap[string(s[i])]
				for !ok {
					i++
					_, ok = tmap[string(s[i])]
				}
			}

			allLessThanZero := true
			for _, v := range tmap {
				if v != 0 {
					allLessThanZero = false
				}
			}

			if allLessThanZero {
				if min == "" || len(min) > len(s[i:j+1]) {
					min = s[i : j+1]
				}
				tmap[string(s[i])]++

				// move i forward until a character in tmap is seen
				i++
				if i < len(s) {
					_, ok := tmap[string(s[i])]
					for !ok {
						i++
						_, ok = tmap[string(s[i])]
					}
				}
			}
		}

		j++
	}

	return min
}
