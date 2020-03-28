package pacific_atlantic_water_flow_417

import "math"

// Second solution: From all locations where water flows in (pacific and atlantic coasts),
// see how far the water would flow in the land and mark those locations as reachable
// by both pacific and atlantic coasts. Then at the end, locations which were reachable
// by both coasts will be added to the final set of coordinates to return.
func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	// create matrices to represent where water can flow
	// into the land from the pacific and the atlantic
	pacific := make([][]bool, len(matrix))
	for i := range pacific {
		pacific[i] = make([]bool, len(matrix[0]))
	}
	atlantic := make([][]bool, len(matrix))
	for i := range atlantic {
		atlantic[i] = make([]bool, len(matrix[0]))
	}

	// explore from the east (atlantic) and west (pacific),
	// marking each spot the water could flow into the land.
	for r := 0; r < len(matrix); r++ {
		explore(matrix, &pacific, math.MinInt32, r, 0)
		explore(matrix, &atlantic, math.MinInt32, r, len(matrix[0])-1)
	}

	// explore from the south (atlantic) and north (pacific),
	// marking each spot the water could flow into the land.
	for c := 0; c < len(matrix[0]); c++ {
		explore(matrix, &pacific, math.MinInt32, 0, c)
		explore(matrix, &atlantic, math.MinInt32, len(matrix)-1, c)
	}

	// coordinates that had atlantic and pacific ocean flow into them are returned
	coordinates := make([][]int, 0)
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if pacific[r][c] && atlantic[r][c] {
				coordinates = append(coordinates, []int{r, c})
			}
		}
	}

	return coordinates
}

func explore(matrix [][]int, visited *[][]bool, prev, r, c int) {
	v := *visited

	// if we can't travel further, then return
	if r < 0 || c < 0 ||
		r >= len(matrix) || c >= len(matrix[0]) ||
		v[r][c] || matrix[r][c] < prev {
		return
	}

	v[r][c] = true
	explore(matrix, visited, matrix[r][c], r-1, c) // north
	explore(matrix, visited, matrix[r][c], r+1, c) // south
	explore(matrix, visited, matrix[r][c], r, c+1) // east
	explore(matrix, visited, matrix[r][c], r, c-1) // west
}

// Initial solution: for each coordinate, see if you can reach
// both the pacific and the atlantic. Correct solution, but is slow.
func pacificAtlantic0(matrix [][]int) [][]int {
	coordinates := make([][]int, 0)
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if canReachPacific(matrix, r, c, r, c) && canReachAtlantic(matrix, r, c, r, c) {
				coordinates = append(coordinates, []int{r, c})
			}
		}
	}

	return coordinates
}

func canReachPacific(matrix [][]int, r, c, pr, pc int) bool {
	if reachedAtlantic(matrix, r, c) {
		return false
	}

	if reachedPacific(r, c) {
		return true
	}

	// go north
	if waterCanFlow(matrix, pr, pc, r, c, r-1, c) &&
		canReachPacific(matrix, r-1, c, r, c) {
		return true
	}

	// go south
	if waterCanFlow(matrix, pr, pc, r, c, r+1, c) &&
		canReachPacific(matrix, r+1, c, r, c) {
		return true
	}

	// go east
	if waterCanFlow(matrix, pr, pc, r, c, r, c+1) &&
		canReachPacific(matrix, r, c+1, r, c) {
		return true
	}

	// go west
	if waterCanFlow(matrix, pr, pc, r, c, r, c-1) &&
		canReachPacific(matrix, r, c-1, r, c) {
		return true
	}

	return false
}

func canReachAtlantic(matrix [][]int, r, c, pr, pc int) bool {
	if reachedPacific(r, c) {
		return false
	}

	if reachedAtlantic(matrix, r, c) {
		return true
	}

	// go north
	if waterCanFlow(matrix, pr, pc, r, c, r-1, c) &&
		canReachAtlantic(matrix, r-1, c, r, c) {
		return true
	}

	// go south
	if waterCanFlow(matrix, pr, pc, r, c, r+1, c) &&
		canReachAtlantic(matrix, r+1, c, r, c) {
		return true
	}

	// go east
	if waterCanFlow(matrix, pr, pc, r, c, r, c+1) &&
		canReachAtlantic(matrix, r, c+1, r, c) {
		return true
	}

	// go west
	if waterCanFlow(matrix, pr, pc, r, c, r, c-1) &&
		canReachAtlantic(matrix, r, c-1, r, c) {
		return true
	}

	return false
}

func reachedAtlantic(matrix [][]int, r, c int) bool {
	return r >= len(matrix) || (r > -1 && c >= len(matrix[r]))
}

func reachedPacific(r, c int) bool {
	return r < 0 || c < 0
}

func waterCanFlow(matrix [][]int, pr, pc, ro, co, rd, cd int) bool {
	// if the next flow destination has reached pacific, atlantic,
	// or has a value less than or equal to the origin value
	return (pr != rd || pc != cd) &&
		(reachedAtlantic(matrix, rd, cd) ||
			reachedPacific(rd, cd) ||
			matrix[ro][co] >= matrix[rd][cd])
}
