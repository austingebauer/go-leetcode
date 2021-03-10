package plus_one_66

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		incr := digits[i] + carry

		if incr == 10 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] = incr
			carry = 0
		}
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
