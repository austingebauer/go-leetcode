package number_of_1_bits_191

func hammingWeight(num uint32) int {
	ones := 0
	for num != 0 {
		if num%2 == 1 {
			ones++
		}

		num = num / 2
	}

	return ones
}
