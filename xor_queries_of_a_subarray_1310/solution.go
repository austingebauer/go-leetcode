package xor_queries_of_a_subarray_1310

func xorQueries(arr []int, queries [][]int) []int {
	res := make([]int, len(queries))
	for idx, q := range queries {
		j := q[0]
		k := q[1]
		if j == k {
			res[idx] = arr[j]
			continue
		}

		val := 0
		for j <= k {
			val ^= arr[j]
			j++
		}
		res[idx] = val
	}

	return res
}
