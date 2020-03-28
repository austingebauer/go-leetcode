package set_matrix_zeros_73

import "math"

func setZeroes(matrix [][]int) {
	// Decide to use the first value of every row and column as the marker
	// for whether the row or column should be set to 0 or not.
	// Since matrix[0][0] can indicate that both 0-th row and column should
	// be set to 0, we need a way to distinguish whether just the 0-th
	// row and column should be set to 0. So, use matrix[0][0] to indicate
	// if row 0 should be set to all zeros, and use an additional variable,
	// setColZero to indicate if col 0 should be set to all zeros.

	setColZero := false
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			v := matrix[r][c]

			// prevent any column 0 values in the rows from marking matrix[0][0]
			// to 0, which indicates row 0 being marked for 0
			if v == 0 && c == 0 {
				setColZero = true
				continue
			}

			if v == 0 {
				// set the row marker to 0
				matrix[r][0] = 0

				// set the column marker to 0
				matrix[0][c] = 0
			}
		}
	}

	// mark cell to 0 if it's row or column marker is 0, starting
	// at matrix[1][1] and using the markers set earlier
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// mark row 0 with 0's
	if matrix[0][0] == 0 {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[0][c] = 0
		}
	}

	// mark col 0 with 0's
	if setColZero {
		for r := 0; r < len(matrix); r++ {
			matrix[r][0] = 0
		}
	}
}

// Note: this was accepted into leetcode and uses O(1) space.
// There is a different solution that was hinted at that works
// better with ALL input. This algorithm could have problems when
// values are equal to math.MaxInt64 and math.MinInt64.
func setZeroes1(matrix [][]int) {
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			v := matrix[r][c]

			// if we see a new 0 or a 0 that was converted into a math.MinInt64
			if v == 0 || v == math.MinInt64 {
				matrix[r][c] = math.MaxInt64

				// set row 0 values
				for s := 0; s < len(matrix[r]); s++ {
					if matrix[r][s] != 0 && matrix[r][s] != math.MinInt64 {
						matrix[r][s] = math.MaxInt64
					} else if matrix[r][s] == 0 {
						matrix[r][s] = math.MinInt64
					}
				}

				// set col 0 values
				for s := 0; s < len(matrix); s++ {
					if matrix[s][c] != 0 && matrix[s][c] != math.MinInt64 {
						matrix[s][c] = math.MaxInt64
					} else if matrix[s][c] == 0 {
						matrix[s][c] = math.MinInt64
					}
				}
			}
		}
	}

	// scan and replace all math.MaxInt32 values into 0 values
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			v := matrix[r][c]
			if v == math.MaxInt64 {
				matrix[r][c] = 0
			}
		}
	}
}
