package spiral_matrix_54

func spiralOrder(matrix [][]int) []int {
	seen := make(map[int]bool)
	spiral := make([]int, 0)
	seenStep := false
	dir := "east"
	var r, c int
	for !seenStep && !isOutOfBounds(matrix, r, c) {
		seen[r*10+c] = true
		spiral = append(spiral, matrix[r][c])

		switch dir {
		case "north":
			r--
		case "east":
			c++
		case "south":
			r++
		case "west":
			c--
		}

		_, seenStep = seen[r*10+c]
		if isOutOfBounds(matrix, r, c) || seenStep {
			switch dir {
			case "north":
				r++
				dir = "east"
				c++
			case "east":
				c--
				dir = "south"
				r++
			case "south":
				r--
				dir = "west"
				c--
			case "west":
				c++
				dir = "north"
				r--
			}
		}

		_, seenStep = seen[r*10+c]
	}

	return spiral
}

func isOutOfBounds(m [][]int, r, c int) bool {
	return r < 0 || r >= len(m) ||
		c < 0 || c >= len(m[r])
}
