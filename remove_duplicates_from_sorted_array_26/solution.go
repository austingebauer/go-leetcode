package remove_duplicates_from_sorted_array_26

func removeDuplicates(nums []int) int {
	i, j, u := 0, 0, 1
	for j < len(nums)-1 {
		j = i
		for j < len(nums)-1 && nums[i] == nums[j] {
			j++
		}
		if nums[i] != nums[j] {
			u++
		}

		for k := j - 1; k > i; k-- {
			nums[k] = nums[j]
		}

		i++
	}

	return u
}
