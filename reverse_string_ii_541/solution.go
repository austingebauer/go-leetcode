package reverse_string_ii_541

import "math"

func reverseStr(s string, k int) string {
	r := []rune(s)
	for i := 0; i < len(s); i += 2 * k {
		// get the index to the end of k chars starting at i
		// pick the minimum for the case where the next k
		// characters falls off the end of s.
		j := int(math.Min(float64(i+k-1), float64(len(s)-1)))

		// reverse the string in the range of [i, j]
		s := i
		for s < j {
			hold := r[s]
			r[s] = r[j]
			r[j] = hold
			s++
			j--
		}
	}

	return string(r)
}

// Note: original solution from mock interview.
// I learned how to reverse string in place by using []rune or []byte.
// Having to reverse by string concatenation make this problem tough.
// Also, the problem description was vague in cases with characters
// remaining to sort.
func reverseStrOrig(s string, k int) string {
	if len(s) < k {
		return s
	}

	i := 0
	for i+k < len(s) {
		end := i + k
		j := end - 1
		revStr := ""
		for j >= i {
			revStr += string(s[j])
			j--
		}

		s = s[0:i] + revStr + s[end:]
		i = i + (2 * k)

		// if there are less than k characters left, reverse all of them
		charsLeft := len(s) - end
		if charsLeft < k {
			// reverse all of them
			e := len(s) - 1
			revStr := ""
			for e >= end {
				revStr += string(s[e])
				e--
			}

			s = s[0:end] + revStr
		}

		// if there are less than 2k but greater than or equal to k characters,
		// then reverse the first k characters and left the other as original
		if charsLeft < 2*k && charsLeft > k {
			// reverse the first k characters
			h := end + k
			e := h - 1
			revStr := ""
			for e >= end {
				revStr += string(s[e])
				e--
			}

			s = s[0:end] + revStr + s[h:]
		}
	}

	return s
}
