package find_the_difference_of_two_arrays_2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := sliceToSet(nums1)
	set2 := sliceToSet(nums2)
	diff := make([][]int, 2)
	diff[0] = setDifference(set1, set2)
	diff[1] = setDifference(set2, set1)
	return diff
}

func setDifference(s1, s2 map[int]struct{}) []int {
	var res []int
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			res = append(res, k)
		}
	}

	return res
}

func sliceToSet(a []int) map[int]struct{} {
	s := make(map[int]struct{})
	for _, v := range a {
		s[v] = struct{}{}
	}
	return s
}
