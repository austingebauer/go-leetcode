package container_with_most_water_11

import "math"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		currentMax := (right - left) *
			int(math.Min(float64(height[left]), float64(height[right])))
		maxArea = int(math.Max(float64(maxArea), float64(currentMax)))

		// We have two lines. The area is bounded by the shorter of the two.
		// If we increase the pointer to the larger of the two, no larger and
		// subsequent line will help maximize the area. If we keep the larger
		// and increase the pointer to the smaller of the two, we stand the
		// chance to maximize the area a bit more when subsequent lines are larger.
		// So, incr/decr the line with the shorter height of the two.
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
