package search_in_rotated_sorted_array_33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// if no rotation happened to the sorted nums, the
	// just binary search over the entire nums
	if nums[0] <= nums[len(nums)-1] {
		// no rotation happened at all
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	// binary search for the inflection point
	start := 0
	end := len(nums) - 1
	var mid int
	for start < end {
		mid = start + ((end - start) / 2)

		// if mid is smallest at inflection
		if mid-1 > -1 && nums[mid] < nums[mid-1] {
			// bring mid back to largest
			mid--
			break
		}

		// if mid is largest at inflection
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			break
		}

		if nums[0] < nums[mid] {
			// inflection point of right of mid
			start = mid
		} else {
			// inflection point of left of mid
			end = mid
		}
	}

	// if the target is less than nums at 0, then it falls on
	// the right-hand side of the inflection point
	if target < nums[0] {
		return binarySearch(nums, target, mid+1, len(nums)-1)
	}

	// if the target is greater than nums at the last index, then
	// it falls on the left-hand side of the inflection point
	if target > nums[len(nums)-1] {
		// target is on left-hand side of inflection point
		return binarySearch(nums, target, 0, mid)
	}

	return -1
}

func binarySearch(nums []int, target, start, end int) int {
	for start <= end {
		mid := start + ((end - start) / 2)
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

// Note: Study again
func search0(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	pivot := findPivot(nums)
	if pivot == 0 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	// If target less than nums[0], it's in nums[pivot:]
	// Ex: 4 5 6 7 |0 1 2|
	if target < nums[0] {
		return binarySearch(nums, pivot, len(nums)-1, target)
	}

	// If target greater than nums[0], it's in nums[:pivot]
	// Ex: |4 5 6 7| 0 1 2
	if target > nums[len(nums)-1] {
		return binarySearch(nums, 0, pivot-1, target)
	}

	return -1
}

func findPivot(nums []int) int {
	left := 0
	right := len(nums) - 1

	// no pivot and rotation happened
	if nums[left] < nums[right] {
		return 0
	}

	for left <= right {
		pivot := (left + right) / 2

		// found the pivot
		if nums[pivot] > nums[pivot+1] {
			return pivot + 1
		}

		// less than or equal
		if nums[pivot] >= nums[left] {
			left = pivot + 1
		} else { // nums[pivot] < nums[left]
			right = pivot - 1
		}
	}

	return 0
}

func binarySearch0(nums []int, left, right, target int) int {
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if target > nums[middle] {
			left = middle + 1
		} else { // target < nums[middle]
			right = middle - 1
		}
	}

	return -1
}
