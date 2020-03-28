package roman_to_integer_13

func romanToInt(s string) int {
	out := 0
	numMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	cIdx := 0
	for cIdx < len(s) {
		c := string(s[cIdx])
		if c == "I" {
			if cIdx < len(s)-1 {
				next := string(s[cIdx+1])
				if next == "V" || next == "X" {
					out += numMap[next] - numMap[c]
					cIdx += 2
					continue
				}
			}
		}
		if c == "X" {
			if cIdx < len(s)-1 {
				next := string(s[cIdx+1])
				if next == "L" || next == "C" {
					out += numMap[next] - numMap[c]
					cIdx += 2
					continue
				}
			}
		}
		if c == "C" {
			if cIdx < len(s)-1 {
				next := string(s[cIdx+1])
				if next == "D" || next == "M" {
					out += numMap[next] - numMap[c]
					cIdx += 2
					continue
				}
			}
		}

		out += numMap[c]
		cIdx++
	}

	return out
}
