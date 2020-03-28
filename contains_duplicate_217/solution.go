package contains_duplicate_217

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return true
		}

		m[n] = true
	}

	return false
}

/*
Note: bit shifting technique only worked for n > -1 in nums
func containsDuplicate(nums []int) bool {
	var dups uint64
	for _, n := range nums {
		exists := dups & (1 << uint64(n))
		if exists == 0 {
			dups |= 1 << uint64(n)
		} else {
			return true
		}
	}

	return false
}
*/
