package two_sum_1

// twoSum does not assume that nums is sorted.
func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return []int{}
	}

	numToIndex := make(map[int]int)
	for idx, num := range nums {
		need := target - num
		mIdx, ok := numToIndex[need]
		if ok {
			return []int{mIdx, idx}
		}

		numToIndex[num] = idx
	}

	return []int{}
}

// twoSumSortedInput assumes that nums is sorted.
func twoSumSortedInput(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return []int{}
	}

	front := 0
	rear := numsLen - 1
	for front != rear {
		twoSum := nums[front] + nums[rear]
		if twoSum == target {
			return []int{
				front,
				rear,
			}
		}

		if twoSum < target {
			front++
		} else {
			rear--
		}
	}

	return []int{}
}
