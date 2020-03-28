package longest_substring_without_repeating_characters_3

import "math"

func lengthOfLongestSubstring(s string) int {
	front := 0
	rear := 0
	longest := 0
	set := make(map[string]bool)
	for front < len(s) {
		_, ok := set[string(s[front])]
		if ok {
			for s[rear] != s[front] {
				delete(set, string(s[rear]))
				rear++
			}
			rear++
		} else {
			set[string(s[front])] = true
			longest = int(math.Max(float64(longest), float64(front-rear+1)))
		}

		front++
	}

	return longest
}
