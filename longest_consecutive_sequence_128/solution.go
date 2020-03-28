package longest_consecutive_sequence_128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Build disjoint set
	djs := make(map[int]int)
	for _, n := range nums {
		djs[n] = n
	}

	// For each kv pair in disjoint set, point it to it's parent if parent exists
	// Parent means it's value plus one
	for k := range djs {
		_, ok := djs[k+1]
		if ok {
			djs[k] = k + 1
		}
	}

	// Traverse disjoint set up-links and count longest seen
	longest := 1
	for k, v := range djs {
		kk := k
		vv := v
		cnt := 1
		for kk != vv {
			cnt++
			kk = vv
			vv = djs[kk]
		}

		if cnt > longest {
			longest = cnt
		}
	}

	return longest
}
