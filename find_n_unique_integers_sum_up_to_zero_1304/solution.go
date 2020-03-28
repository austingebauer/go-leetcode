package find_n_unique_integers_sum_up_to_zero_1304

// Time: O(n)
// Space: O(n) (output space)
func sumZero(n int) []int {
	uniques := make([]int, 0)
	var neg, pos = -1, 1
	for i := 0; i < n/2; i++ {
		uniques = append(uniques, []int{pos, neg}...)
		neg--
		pos++
	}

	if n%2 == 1 {
		uniques = append(uniques, 0)
	}

	return uniques
}

func sumZeroNoAppend(n int) []int {
	uniques := make([]int, n)
	placement := 0
	var neg, pos = -1, 1
	for i := 0; i < n/2; i++ {
		uniques[placement] = pos
		placement++

		uniques[placement] = neg
		placement++

		neg--
		pos++
	}

	if n%2 == 1 {
		uniques[placement] = 0
	}

	return uniques
}
