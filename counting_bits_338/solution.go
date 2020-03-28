package counting_bits_338

func countBits(num int) []int {
	bits := make([]int, num+1)
	for n := 1; n <= num; n++ {
		bits[n] = bits[n&(n-1)] + 1
	}

	return bits
}
