package palindrome_number_9

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := x

	// count the number of digits in x
	digits := 0
	for n != 0 {
		n = n / 10
		digits++
	}

	// for each digits / 2 times, compare left hand side digits to right hand side digits
	place := 0
	for i := 1; i <= digits/2; i++ {
		hold := x
		ld := (hold / int(math.Pow(float64(10), float64(place)))) % 10

		for j := 0; j < digits-i; j++ {
			hold = hold / 10
		}

		if ld != hold%10 {
			return false
		}

		place++
	}

	return true
}
