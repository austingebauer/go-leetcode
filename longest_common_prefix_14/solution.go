package longest_common_prefix_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// find min len string in strs
	idx := 0
	minLen := len(strs[idx])
	for i, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
			idx = i
		}
	}

	// for each char in min len string,
	// see if every nth char in strs matches
	// to build the longest common prefix.
	lPrefix := ""
	for i := 0; i < minLen; i++ {
		allMatch := true
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[idx][i] {
				allMatch = false
			}
		}

		if allMatch {
			lPrefix += string(strs[idx][i])
		} else {
			break
		}
	}

	return lPrefix
}
