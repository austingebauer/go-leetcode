package convert_integer_to_the_sum_of_two_no_zero_integers

func getNoZeroIntegers(n int) []int {
	one := n
	two := 0

	for hasZeros(one) || hasZeros(two) {
		one--
		two++
	}

	return []int{two, one}
}

func hasZeros(n int) bool {
	if n == 0 {
		return true
	}

	for n != 0 {
		if n%10 == 0 {
			return true
		}

		n /= 10
	}

	return false
}
