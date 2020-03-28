package reverse_words_in_a_string_151

import "strings"

// Note: solution using a single string with which tokens are prepended into.
func reverseWords(s string) string {
	s = strings.Trim(s, " ")

	sFinal := ""
	start := 0
	finish := 0
	for finish < len(s) {
		if string(s[finish]) == " " {
			// prepend the token found into the final string
			// this is slow, "don't concat strings in a loop"
			sFinal = s[start:finish] + " " + sFinal

			// clear whitespace in between tokens for next start
			for string(s[finish]) == " " {
				finish++
			}

			start = finish
		} else {
			finish++
		}
	}

	// Prepend final token as finish stopped at the end of s
	sFinal = s[start:finish] + " " + sFinal

	return strings.TrimSuffix(sFinal, " ")
}

// Note: First, correct, and optimal runtime solution using an array
// for tokens and reversing the array.
func reverseWordsArrayReversal(s string) string {
	s = strings.Trim(s, " ")

	tokens := make([]string, 0)
	start := 0
	finish := 0
	for finish < len(s) {
		if string(s[finish]) == " " {
			// is the end of the token
			tokens = append(tokens, s[start:finish])

			// clear whitespace in between tokens for next start
			for string(s[finish]) == " " {
				finish++
			}

			start = finish
		} else {
			finish++
		}
	}
	// Add final token as finish stopped at the end of s
	tokens = append(tokens, s[start:finish])

	// Reverse the slice of tokens
	reverse(tokens)

	// Join each token on a single space
	return strings.Join(tokens, " ")
}

func reverse(a []string) {
	front := 0
	back := len(a) - 1
	for front < back {
		hold := a[back]
		a[back] = a[front]
		a[front] = hold

		front++
		back--
	}
}
