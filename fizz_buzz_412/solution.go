package fizz_buzz_412

import "strconv"

func fizzBuzz(n int) []string {
	a := make([]string, n)
	for i := 1; i <= n; i++ {
		fizz := ""
		buzz := ""

		if i%3 == 0 {
			fizz = "Fizz"
		}

		if i%5 == 0 {
			buzz = "Buzz"
		}

		r := fizz + buzz
		if r == "" {
			a[i-1] = strconv.Itoa(i)
		} else {
			a[i-1] = r
		}
	}

	return a
}
