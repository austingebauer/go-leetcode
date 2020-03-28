package count_and_say_38

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	sNext := ""
	ch := string(s[0])
	freq := 0
	for _, c := range s {
		if string(c) == ch {
			freq++
		} else {
			sNext += fmt.Sprintf("%d%s", freq, ch)
			freq = 1
			ch = string(c)
		}
	}

	sNext += fmt.Sprintf("%d%s", freq, ch)
	return sNext
}
