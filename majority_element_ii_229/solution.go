package majority_element_ii_229

// Note: study again
func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// At most, nums can contain 2 majority elements which
	// occur more than len(nums) / 3 times. So, we need two
	// counts and two candidates for the Moore Voting algorithm.
	count1, count2 := 0, 0

	// Pick candidates (they is arbitrary)
	candidate1, candidate2 := 10, 11

	for _, num := range nums {
		if num == candidate1 {
			// Saw num that is equal to candidate 1, increase it's count
			count1++
		} else if num == candidate2 {
			// Saw num that is equal to candidate 2, increase it's count
			count2++
		} else if count1 == 0 {
			// Replacing first candidate
			candidate1, count1 = num, 1
		} else if count2 == 0 {
			// Replacing second candidate
			candidate2, count2 = num, 1
		} else {
			// Both candidates counts decrease
			count1--
			count2--
		}
	}

	// At this point, candidate1 and candidate2 are set to majority elements in nums
	// Tally how many times they were seen in nums bc the algorithm doesn't
	// guarantee a majority element.
	count1 = 0
	count2 = 0
	for _, n := range nums {
		if n == candidate1 {
			count1++
		}

		if n == candidate2 {
			count2++
		}
	}

	// if counts for candidate1 and candidate2 are larger than len(nums)/3, return them
	maj := make([]int, 0)
	if count1 > len(nums)/3 {
		maj = append(maj, candidate1)
	}

	if count2 > len(nums)/3 {
		maj = append(maj, candidate2)
	}

	return maj
}
