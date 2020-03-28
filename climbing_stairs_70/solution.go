package climbing_stairs_70

func climbStairs(n int) int {
	memo := make([]int, n)
	return exploreStairs(n, 0, memo)
}

func exploreStairs(n int, c int, memo []int) int {
	if c > n {
		return 0
	}

	if c == n {
		return 1
	}

	if memo[c] > 0 {
		return memo[c]
	}

	memo[c] = exploreStairs(n, c+1, memo) + exploreStairs(n, c+2, memo)
	return memo[c]
}
