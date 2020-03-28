package fibonacci_number_509

func fib(N int) int {
	if N < 2 {
		return N
	}

	return fib(N-1) + fib(N-2)
}
