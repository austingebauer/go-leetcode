package sum_of_two_integers_371

// Note: study again. Cool problem!
func getSum(a int, b int) int {
	for b != 0 {
		// corresponding bits (1, 1) will need to be set to 0 and carry a 1 over
		// carry holds the slots where we'll need to carry
		carry := a & b

		// corresponding bits (1, 0) and (0, 1) will stay set
		// a no longer has locations of carry bits set via XOR
		a = a ^ b

		// set b to the result of carrying over each location that we need to carry in
		// eventually all carries will find a slot in a (may take many carries),
		// after which b will become 0 and the loop condition will be false.
		b = carry << 1
	}

	return a
}
