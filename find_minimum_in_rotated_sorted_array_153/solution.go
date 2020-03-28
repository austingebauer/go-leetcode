package find_minimum_in_rotated_sorted_array_153

// Time: O(log(n))
// Space: O(1)
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// We can use a binary search to find the 'inflection point',
	// which is the point where nums[i] > nums[i+1]
	start := 0
	end := len(nums) - 1

	// If the num at start is less than the number at
	// the end, then the array has not been rotated
	if nums[0] < nums[end] {
		return nums[0]
	}

	// Binary search over nums for the 'inflection point'
	// Since we don't know where this point is, we can make
	// a decision and see **what would get use closer to it**.
	// Observe that if the number at the beginning of nums
	// is larger than the mid, then this point is on the left
	// side of mid. If the number at the beginning of nums is less
	// than the mid, then this point is on the right side of mid.
	min := nums[0]
	for start != end {
		// pick a middle index
		midIdx := start + ((end - start) / 2)

		// check if we've found the inflection point, where
		// nums[midIdx] could be less than nums[midIdx-1]
		if midIdx-1 > -1 && nums[midIdx] < nums[midIdx-1] {
			return nums[midIdx]
		}

		// check if we've found the inflection point, where
		// nums[midIdx] could be greater than nums[midIdx+1]
		if midIdx+1 < len(nums) && nums[midIdx] > nums[midIdx+1] {
			return nums[midIdx+1]
		}

		// close is on the search space
		if nums[midIdx] > nums[0] {
			// inflection point is on the right side of mid
			start = midIdx
		} else {
			// nums[midIdx] < nums[0], so inflection point
			// is on the left side of mid
			end = midIdx
		}
	}

	return min
}

// Time: O(n)
// Space: O(1)
func findMinLinear(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			min = nums[i]
			break
		}
	}

	return min
}
