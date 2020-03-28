package decrypt_string_from_alphabet_to_integer_mapping_1309

import "strconv"

func freqAlphabets(s string) string {
	a := "abcdefghijklmnopqrstuvwxyz"
	d := ""
	i := 0
	for i < len(s) {
		if i < len(s)-2 && string(s[i+2]) == "#" {
			alphaIdx, _ := strconv.Atoi(s[i : i+2])
			d += string(a[alphaIdx-1])
			i += 3
		} else {
			alphaIdx, _ := strconv.Atoi(string(s[i]))
			d += string(a[alphaIdx-1])
			i += 1
		}
	}

	return d
}
