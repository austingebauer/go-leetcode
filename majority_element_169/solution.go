package majority_element_169

// Note: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
// Good problem and intro to a new algorithm.
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for i := 0; i < len(nums); i++ {
		// pick first candidate or last suffix brought count to 0
		if count == 0 {
			candidate = nums[i]
		}

		// if nums[i] is equal to the candidate, increase the count by 1
		if nums[i] == candidate {
			count++
		} else {
			// otherwise, decrease the count by 1
			count--
		}
	}

	// if there is a majority, it'll be in candidate because
	// there will be more count++ than count-- for it.
	return candidate
}
