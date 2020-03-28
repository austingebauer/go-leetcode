package partition_to_k_equal_sum_subsets_698

import "sort"

// Note: study again. Good problem!
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	// if we can't evenly distribute the sum of each num over k partitions
	if sum%k != 0 {
		return false
	}

	// the target sum for each partition
	target := sum / k

	// if the largest value in the array is greater than the target sum, return false
	// ex: [5,2,1] k=2 target=4 (we can't use 5 for any partition)
	sort.Ints(nums)
	end := len(nums) - 1
	if nums[end] > target {
		return false
	}

	// clip off values that would fit directly into a partition and reduce k
	for end >= 0 && nums[end] == target {
		end--
		k--
	}

	return fill(make([]int, k), end, nums, target)
}

func fill(subsets []int, idx int, nums []int, target int) bool {
	if idx < 0 {
		return true
	}

	// choose a value from nums to try to slot into each subset
	pick := nums[idx]
	idx--

	// try to fill each subset
	for i := 0; i < len(subsets); i++ {
		if subsets[i]+pick <= target {
			subsets[i] += pick

			// explore if the subsets current value + chosen value is <= target
			if fill(subsets, idx, nums, target) {
				// if call to fill is true, we were able to partition successfully
				// if not, continue on the search
				return true
			}

			// un-choose the value for the current subset
			subsets[i] -= pick
		}

		// if current subset became 0 from the un-choose, then break
		// this keeps all unfilled subsets on the end and reduces work
		if subsets[i] == 0 {
			break
		}
	}

	return false
}

/*
// First take on the problem without peeking solution.
// Pretty close, but off on a couple of ideas.
func canPartitionKSubsets1(nums []int, k int) bool {
	// if there are not enough numbers to partition into k subsets
	if len(nums) < k {
		return false
	}

	// if there are exactly enough numbers to partition into k subsets,
	// then each number must be equal for sums to be equal
	if len(nums) == k {
		set := make(map[int]bool)
		for _, n := range nums {
			set[n] = true
		}

		return len(set) == 1
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	// the target sum for each partition
	target := sum / k
	return fill(nums, make([][]int, k), target, 0)
}

func fill(nums []int, subsets [][]int, target int, fillIndex int) bool {
	// if we have filled each subset
	if fillIndex == len(subsets) {
		return true
	}

	for i, n := range nums {
		// if n is greater than the target, return false
		if n > target {
			return false
		}

		subsetSum := 0
		for _, v := range subsets[fillIndex] {
			subsetSum += v
		}

		if subsetSum + n == target {
			subsets[fillIndex] = append(subsets[fillIndex], n)
			fillIndex++
		} else if subsetSum + n < target {
			subsets[fillIndex] = append(subsets[fillIndex], n)
		}

		nextChoices := append(nums[0:i], nums[i+1:]...)

		filled := fill(nextChoices, subsets, target, fillIndex)

		if filled {
			return true
		}
	}

	return false
}
*/
