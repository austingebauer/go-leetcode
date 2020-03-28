package three_sum_15

import "sort"

func threeSum(nums []int) [][]int {
	triplets := make([][]int, 0)
	sort.Ints(nums)

	if len(nums) < 3 {
		return triplets
	}

	// pick a number, then solve the two sum problem in the space in front of it
	for i := 0; i < len(nums); i++ {
		// skip duplicates of the starting number if we've seen it in the past
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// determine what is needed from two sum
		need := 0 - nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {
			// if we've found what we need
			if nums[start]+nums[end] == need {
				triplets = append(triplets, []int{nums[i], nums[start], nums[end]})

				// Choose to skip duplicates on one end or the other (start or end)
				// In my first answer below, I skipped duplicates from start, for example.
				// We need to continue to search in the space in between start and end
				// for more sums that equal need. So, we need to move one end forward
				// and skip duplicates in order to satisfy the 'unique' requirement of
				// the problem.
				end--
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			} else if nums[start]+nums[end] > need {
				end--
			} else {
				start++
			}
		}
	}

	return triplets
}

func threeSum0(nums []int) [][]int {
	triplets := make([][]int, 0)
	if len(nums) < 3 {
		return triplets
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			tripletSum := nums[i] + nums[j] + nums[k]
			if tripletSum == 0 {
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				j++

				// We found the sum we wanted, and moved j forward one index.
				// If the value at j is equal to the value at j-1, a duplicate
				// will be produced. Continue to move j forward until a new value
				// is seen or j >= k.
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if tripletSum < 0 {
				// If tripletSum is less than 0, we must increase the sum
				j++
			} else {
				// If tripletSum is less than 0, we must decrease the sum
				k--
			}
		}
	}

	return triplets
}
