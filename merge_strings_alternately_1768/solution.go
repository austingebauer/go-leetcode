package add_two_numbers_2

func mergeAlternately(word1 string, word2 string) string {
	var (
		res     string
		i, j, k int
	)

	for i+j < len(word1)+len(word2) {
		if i < len(word1) && k%2 == 0 {
			res += string(word1[i])
			i++
			if j < len(word2) {
				k++
			}
		}
		if j < len(word2) && k%2 == 1 {
			res += string(word2[j])
			j++
			if i < len(word1) {
				k++
			}
		}
	}

	return res
}
