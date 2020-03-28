package reorder_data_in_log_files_937

import (
	"regexp"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	// insertion sort with custom comparison function
	for i := 1; i < len(logs); i++ {
		k := i
		for k > 0 {
			ct := compareLogs(logs[k-1], logs[k])

			if ct <= 0 {
				break
			}

			if ct > 0 {
				hold := logs[k]
				logs[k] = logs[k-1]
				logs[k-1] = hold
			}

			k--
		}
	}

	return logs
}

func compareLogs(a, b string) int {
	aTokens := strings.Split(a, " ")
	bTokens := strings.Split(b, " ")
	aIsDigit, _ := regexp.MatchString("[0-9]", aTokens[1])
	bIsDigit, _ := regexp.MatchString("[0-9]", bTokens[1])

	// don't move anything if a and b are digits
	// don't move back b if it's a digit and a is not
	if (aIsDigit && bIsDigit) || (!aIsDigit && bIsDigit) {
		return -1
	}

	// move back b if a is a digit and b is not
	if aIsDigit && !bIsDigit {
		return 1
	}

	// both a & b aren't digit logs
	aLetters := strings.TrimPrefix(a, aTokens[0]+" ")
	bLetters := strings.TrimPrefix(b, bTokens[0]+" ")
	if aLetters == bLetters {
		// if aLetter and bLetters are equal, compare identifiers
		return strings.Compare(aTokens[0], bTokens[0])
	}

	// compare letters
	return strings.Compare(aLetters, bLetters)
}
