package merge_sorted_array_88

// O(m+n) solution
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	n1, n2, p := m-1, n-1, m+n-1
	for n2 > -1 {
		if n1 < 0 || nums2[n2] >= nums1[n1] {
			nums1[p] = nums2[n2]
			n2--
		} else {
			nums1[p] = nums1[n1]
			n1--
		}
		p--
	}
}

// Note: solution from rotate_string_796 interview that stumped me for a bit.
// Wanted to solve without shifting and was tricky.
func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	// place largest values into the back of nums1
	p := len(nums1) - 1
	i := m - 1
	j := n - 1

	// while i and j are greater than 0
	for i >= 0 && j >= 0 {
		// place largest of i and j into p
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}

		p--
	}

	// if j is not less than 0, fill nums1 with num2 values
	for j >= 0 {
		nums1[j] = nums2[j]
		j--
	}
}

// Note: original solution with array shifting
func merge0(nums1 []int, m int, nums2 []int, n int) {
	slot := 0

	// for each in nums2, insert into nums1
	for i := 0; i < n; i++ {

		// find the slot to insert nums2[i] into
		for slot < m+i && nums1[slot] <= nums2[i] {
			slot++
		}

		// count how many shifts are needed
		shiftCnt := m - slot + 1 + i

		// shift over each number to insert nums2[i]
		for s := 0; s < shiftCnt; s++ {
			hold := nums2[i]
			nums2[i] = nums1[slot+s]
			nums1[slot+s] = hold
		}
	}
}
