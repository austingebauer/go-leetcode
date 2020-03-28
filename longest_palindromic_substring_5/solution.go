package longest_palindromic_substring_5

import "math"

// expand at the center 2*n - 1 times for
// palindromes with even and odd lengths
func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		s1 := expandPalindrome(s, i, i)
		s2 := expandPalindrome(s, i, i+1)

		if len(s1) > len(longest) {
			longest = s1
		}
		if len(s2) > len(longest) {
			longest = s2
		}
	}

	return longest
}

func expandPalindrome(s string, i, j int) string {
	if i < 0 || j >= len(s) {
		return ""
	}

	// Don't expand if they're not already equal
	if s[i] != s[j] {
		return ""
	}

	// expand i and j until they're out-of-bounds
	for i > -1 && j < len(s) {
		// if expanding would make the string no longer
		// a palindrome or go out-of-bounds, break so
		// we can return the palindromic substring
		if i-1 < 0 || j+1 >= len(s) || s[i-1] != s[j+1] {
			break
		}

		i--
		j++
	}

	// return the palindromic substring
	return s[i : j+1]
}

// Note: study again
func longestPalindrome0(s string) string {
	if len(s) < 1 {
		return ""
	}

	// For each index in s, expand as much as possible at every center
	// in the string. There are n * 2 - 1 centers of possible palindromes in s.
	// That is true because there is a center at every index i and a center
	// in between every two indices i and j.
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := int(math.Max(float64(len1), float64(len2)))

		// If the maxLen from the last iteration is longer in length
		// than the longest palindromic substring seen so far.
		if maxLen > end-start {
			// Calculate the starting index of the longest substring so far
			// by:
			//   1. subtract 1 from the maxLen that was expanded from the center
			//   2. divide that value by 2
			//   3. subtract that value from i

			// Example:
			//   01234
			//   babad
			//     i

			//   Minus one on maxLen for even number palindromes
			//   0123
			//   abba

			//   i = 2, maxLen = 3
			//   start = i - ((maxLen - 1) / 2)
			//   start = 2 - ((3 - 1) / 2)
			//   start = 1
			start = i - ((maxLen - 1) / 2)
			end = i + (maxLen / 2)
		}
	}

	// return the substring from start to end (+1 for non-inclusive substring)
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	// The left and right values are either:
	//   - equal
	//   - right = left + 1

	// while left is greater than -1 and
	// right is less than the length of s and
	// s[left] is equal to s[right].
	for left > -1 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	// return the length of the palindromic expansion (i.e., units from
	// center being the passed left and right) about the center.
	return right - left - 1
}
