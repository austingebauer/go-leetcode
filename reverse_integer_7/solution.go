package reverse_integer_7

import "math"

func reverse(x int) int {
	var out int64
	for x != 0 {
		out = (out * 10) + int64(x%10)
		x = x / 10
	}

	// if out overflows positive (2^32 - 1) or negative (-2^31) 32bit integer
	if out > math.MaxInt32 || out < math.MinInt32 {
		return 0
	}

	return int(out)
}
