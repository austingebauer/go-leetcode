package valid_parentheses_20

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	stack := make([]string, len(s))
	stackIdx := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])

		// is an open character, so push on stack
		if c == "(" || c == "{" || c == "[" {
			stack[stackIdx] = c
			stackIdx++
		} else {
			// initial closing character special case
			if stackIdx == 0 {
				return false
			}

			// is closing char, so pop one off stack to compare
			stackIdx--
			popC := stack[stackIdx]

			if (c == ")" && popC != "(") ||
				(c == "}" && popC != "{") ||
				(c == "]" && popC != "[") {
				return false
			}
		}
	}

	return stackIdx == 0
}
