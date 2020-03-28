package maximum_length_of_a_concatenated_string_with_unique_characters_1239

import "math"

// Note: Study again. Good problem with bit manipulation to detect
// if a string has unique characters.
func maxLength(arr []string) int {
	return explore(arr, 0, 0)
}

func explore(arr []string, index int, usedChar int) int {
	if index == len(arr) {
		return 0
	}

	// does arr[index] have any characters with an occurrence greater than 1?
	currentUsedChar := 0
	s := arr[index]
	valid := true
	for i := 0; i < len(s); i++ {
		// currentlyUsedChar is a bit mask which represents which characters
		// have been used. In the mask, each used character (flipped on by shifting
		// (1 << (s[i] - 'a') the distance from char to 'a'.

		// If there is an overlapping character '00000100' & '00000100', the
		// bitwise AND will *not* return 0, indicating that we've seen a duplicate.
		exists := currentUsedChar & (1 << (s[i] - 'a'))
		if exists == 0 {
			// update currentUsedChar with bitwise OR, which will keep
			// old 1 bits and add new 1 bits to the mask.
			currentUsedChar |= 1 << (s[i] - 'a')
		} else {
			valid = false
			break
		}
	}

	output := explore(arr, index+1, usedChar)
	if valid && (usedChar&currentUsedChar) == 0 {
		output = int(math.Max(float64(output),
			float64(len(s)+explore(arr, index+1, usedChar|currentUsedChar))))
	}

	return output
}

/*
func maxLength(arr []string) int {
	// build a count of frequencies
	f := make([]int, 26)
	for _, a := range arr {
		for i := 0; i < len(a); i++ {
			f[int(a[i])-int('a')]++
		}
	}

	return explore(arr, make(map[string]int), f)
}

func explore(arr []string, memo map[string]int, freq []int) int {
	if len(arr) == 0 {
		return 0
	}

	s := strings.Join(arr, "")

	v, ok := memo[s]
	if ok {
		return v
	}

	// s has unique chars, and s is not going to get larger
	// by exploring more, so return the length of s
	if hasUniqueChars(freq) {
		memo[s] = len(s)
		return len(s)
	}

	maxLen := -1
	for i := range arr {
		// decrement frequencies for chars at arr[i] before exploring
		for j := 0; j < len(arr[i]); j++ {
			freq[int(arr[i][j])-int('a')]--
		}

		// remove the choice at arr[i] and explore
		nextArr := append(append([]string{}, arr[0:i]...), arr[i+1:]...)
		maxLen = int(math.Max(float64(maxLen), float64(explore(nextArr, memo, freq))))

		// increment frequencies for chars at arr[i] after exploring
		for j := 0; j < len(arr[i]); j++ {
			freq[int(arr[i][j])-int('a')]++
		}
	}

	return maxLen
}

func hasUniqueChars(f []int) bool {
	for _, o := range f {
		if o > 1 {
			return false
		}
	}
	return true
}
*/
