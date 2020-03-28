package palindromic_substrings_647

func countSubstrings(s string) int {
	count := 0

	// for each location which can be expanded into a palindrome
	for i := 0; i < len(s); i++ {
		count += expandCenter(s, i, i)
		count += expandCenter(s, i, i+1)
	}

	return count
}

// expandCenter returns the number of palindromic strings
// discovered by expanding around the passed left and right
// index in the passed string s.
func expandCenter(s string, left, right int) int {
	cnt := 0
	for left > -1 && right < len(s) && s[left] == s[right] {
		left--
		right++
		cnt++
	}

	return cnt
}
