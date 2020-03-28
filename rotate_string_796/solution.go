package rotate_string_796

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		return true
	}

	a := A
	for i := 0; i < len(A); i++ {
		if a == B {
			return true
		}

		// rotate a by prepending end to front
		a = a[len(A)-1:] + a[0:len(A)-1]
	}

	return false
}
