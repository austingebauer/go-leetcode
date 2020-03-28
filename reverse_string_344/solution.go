package reverse_string_344

func reverseString(s []byte) {
	for i := 0; i < len(s)-1; i++ {
		j := len(s) - 1

		for i < j {
			hold := s[j]
			s[j] = s[j-1]
			s[j-1] = hold
			j--
		}
	}
}
