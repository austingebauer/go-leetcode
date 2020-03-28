package number_of_islands_200

func numIslands(grid [][]byte) int {
	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				count++
				exploreIsland(grid, r, c)
			}
		}
	}

	return count
}

func exploreIsland(grid [][]byte, r int, c int) {
	grid[r][c] = 0

	// north
	if r-1 > -1 && grid[r-1][c] == 1 {
		exploreIsland(grid, r-1, c)
	}

	// east
	if c+1 < len(grid[r]) && grid[r][c+1] == 1 {
		exploreIsland(grid, r, c+1)
	}

	// south
	if r+1 < len(grid) && grid[r+1][c] == 1 {
		exploreIsland(grid, r+1, c)
	}

	// west
	if c-1 > -1 && grid[r][c-1] == 1 {
		exploreIsland(grid, r, c-1)
	}
}
