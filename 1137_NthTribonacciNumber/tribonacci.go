package tribonacci

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		fib1 := 0
		fib2 := 1
		fib3 := 1
		fib4 := 0
		for i := 3; i <= n; i++ {
			fib4 = fib1 + fib2 + fib3
			fib1 = fib2
			fib2 = fib3
			fib3 = fib4
		}
		return fib4
	}
}
