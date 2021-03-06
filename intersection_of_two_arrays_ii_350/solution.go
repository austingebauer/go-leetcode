package intersection_of_two_arrays_ii_350

func intersect(nums1 []int, nums2 []int) []int {
	counts := make(map[int]int)
	for _, n := range nums1 {
		counts[n]++
	}

	res := make([]int, 0)
	for _, n := range nums2 {
		if v, ok := counts[n]; ok && v > 0 {
			res = append(res, n)
			counts[n]--
		}
	}

	return res
}
