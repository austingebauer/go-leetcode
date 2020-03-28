package happy_number_202

import "math"

func isHappy(n int) bool {
	cycles := false
	seen := make(map[int]bool)
	for n != 1 {
		_, ok := seen[n]
		if ok {
			cycles = true
			break
		}
		seen[n] = true

		sum := 0
		for n != 0 {
			sum += int(math.Pow(float64(n%10), 2))
			n = n / 10
		}
		n = sum
	}

	return !cycles
}
