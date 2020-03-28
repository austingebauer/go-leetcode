package trapping_rain_water_42

import "math"

func trap(height []int) int {
	left := 0
	leftMax := 0
	right := len(height) - 1
	rightMax := 0
	count := 0

	for left < right {
		leftMax = int(math.Max(float64(height[left]), float64(leftMax)))
		rightMax = int(math.Max(float64(height[right]), float64(rightMax)))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		if leftMax > height[left] && leftMax < rightMax {
			count += leftMax - height[left]
		}

		if rightMax > height[right] && leftMax >= rightMax {
			count += rightMax - height[right]
		}
	}

	return count
}

// trap2 worked, but ran out of memory on leetcode.
// This was my first pass at the problem.
// I didn't think clearly enough about the memory, which
// could've been a max(height) x len(height) matrix.
func trap2(height []int) int {
	cols := len(height)
	rows := 0
	for _, h := range height {
		rows = int(math.Max(float64(h), float64(rows)))
	}

	repr := make([][]int, rows)
	for i := range repr {
		repr[i] = make([]int, cols)
	}

	fillCol := 0
	fillRow := rows - 1
	for _, h := range height {
		hDecr := h
		for hDecr > 0 {
			repr[fillRow][fillCol] = 1
			fillRow--
			hDecr--
		}
		fillCol++
		fillRow = rows - 1
	}

	waterCnt := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if repr[row][col] == 1 {
				colCur := col + 1
				colCnt := 0
				foundWall := false
				for colCur < cols {
					if repr[row][colCur] == 1 {
						foundWall = true
						break
					}
					colCur++
					colCnt++
				}

				if foundWall {
					waterCnt += colCnt
				}
			}
		}
	}

	return waterCnt
}
