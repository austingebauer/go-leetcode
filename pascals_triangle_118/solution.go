package pascals_triangle_118

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	for r := 0; r < numRows; r++ {
		row := make([]int, r+1)
		row[0], row[len(row)-1] = 1, 1

		// Build the new row by using values from the previous row
		for c := 1; c < len(row)-1; c++ {
			row[c] = res[r-1][c-1] + res[r-1][c]
		}

		res = append(res, row)
	}

	return res
}
