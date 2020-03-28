package diet_plan_performance_1176

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	start := 0
	end := k - 1
	t := 0
	for i := start; i <= end; i++ {
		t += calories[i]
	}

	points := getPointsForT(t, lower, upper)

	for end < len(calories)-1 {
		t = t - calories[start]
		start++
		end++
		t = t + calories[end]

		points += getPointsForT(t, lower, upper)
	}

	return points
}

func getPointsForT(t, lower, upper int) int {
	if t < lower {
		return -1
	}

	if t > upper {
		return 1
	}

	return 0
}
