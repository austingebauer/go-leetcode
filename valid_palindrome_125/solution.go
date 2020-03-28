package valid_palindrome_125

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	front := 0
	back := len(s) - 1
	s = strings.ToLower(s)

	for front <= back {
		for front < len(s) && !isAlphaNumeric(string(s[front])) {
			front++
		}

		for back > -1 && !isAlphaNumeric(string(s[back])) {
			back--
		}

		// "a:a" case needs initial condition here
		if front <= back && s[front] != s[back] {
			return false
		}

		front++
		back--
	}

	return true
}

func isAlphaNumeric(s string) bool {
	matched, _ := regexp.MatchString("[a-z0-9]", s)
	return matched
}
